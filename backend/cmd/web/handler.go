package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

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

	sessionManager.Put(r.Context(), "user_id", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"message": "Authenticated Successfully",
		"user": map[string]any{
			"id":       id,
			"username": user.Username,
			"email":    email,
		},
	})
}

func (app *application) createRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	userId := sessionManager.Get(r.Context(), "user_id")

	if userId == nil {
		w.WriteHeader(http.StatusUnauthorized)
		response := map[string]string{
			"status":  "error",
			"message": "Unauthorized",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	userIdInt := userId.(int)

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		log.Printf("Decode error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"status":  "error",
			"message": "Bad Request",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if room.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"status":  "error",
			"message": "Room name is required",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if room.MaxPlayers < 2 || room.MaxPlayers > 8 {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"status":  "error",
			"message": "Max players must be between 2 and 8",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if room.IsPrivate && room.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"status":  "error",
			"message": "Password is required for private rooms",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	room_id, err := app.rooms.Create(userIdInt, room.Name, room.MaxPlayers, room.IsPrivate, room.Password)
	if err != nil {
		log.Printf("Create room error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	player := models.PlayerData{
		Username: "test_username",
		ID:       userIdInt,
	}

	room.ID = room_id
	room.CreatedBy = userIdInt
	room.Players = []models.PlayerData{
		player,
	}

	sessionManager.Put(r.Context(), "room_id", room_id)
	activeRooms[room_id] = room

	response := map[string]interface{}{
		"status": "success",
		"room":   room,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
