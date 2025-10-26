package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"playwhot.io/pkg/models"
)

var activeRooms = map[int]models.Room{}
var roomMutex sync.RWMutex

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	response := fmt.Sprintf("{'status':'available','enviroment:%s','version:%s'}", app.environment, app.version)

	w.Write([]byte(response))
}

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Decode error: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = app.users.Insert(user.Username, user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		response := map[string]string{
			"status":  "error",
			"message": err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]string{
		"status":  "success",
		"message": "Account created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// Use json.Decoder with DisallowUnknownFields for faster parsing
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var user models.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Early validation
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}

	id, email, err := app.users.Authenticate(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// Set session first (faster than JSON encoding)
	sessionManager.Put(r.Context(), "user_id", id)

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"message": "Authenticated Successfully",
		"user": map[string]string{
			"username": user.Username,
			"email":    email,
		},
	})
}

func (app *application) createRoom(w http.ResponseWriter, r *http.Request) {
	user_id := sessionManager.Get(r.Context(), "user_id")

	if user_id == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	room_id, err := app.rooms.Create(user_id.(int))
	if err != nil {
		log.Printf("Create room error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sessionManager.Put(r.Context(), "room_id", room_id)
	//TODO: Change this to Redis for scalability
	room := models.Room{
		ID:        room_id,
		CreatedBy: user_id.(int),
		Timestamp: time.Now(),
		Members:   []int{user_id.(int)},
	}
	activeRooms[room_id] = room

	response := map[string]any{
		"status":  "success",
		"message": "Room created successfully",
		"room_id": room_id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (app *application) joinRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	user_id := sessionManager.Get(r.Context(), "user_id")

	if user_id == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&room)

	if err != nil {
		log.Printf("Decode error: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if sessionManager.Exists(r.Context(), "room_id") {
		log.Println("Already belong to a room")
	}

	roomMutex.Lock()
	defer roomMutex.Unlock()

	existingRoom, exists := activeRooms[room.ID]
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	for _, member := range existingRoom.Members {
		if member == user_id.(int) {
			http.Error(w, "Already in room", http.StatusConflict)
			return
		}
	}

	if len(existingRoom.Members) >= 4 {
		http.Error(w, "Room is full", http.StatusConflict)
		return
	}

	existingRoom.Members = append(existingRoom.Members, user_id.(int))
	activeRooms[room.ID] = existingRoom

	sessionManager.Put(r.Context(), "room_id", room.ID)

	response := map[string]string{
		"status":  "success",
		"message": "Joined room successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
