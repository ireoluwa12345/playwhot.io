## Plan: Implementing handleCardPlay Handler

### Phase 1: Message Processing Infrastructure

1. Create message router in cmd/web/handler.go to route WebSocket messages
2. Modify client readPump to call message router instead of blind
broadcasting
3. Add game message handlers for different message types

### Phase 2: handleCardPlay Implementation

1. Create handleCardPlay function that:
 • Parses card data from WebSocket message
 • Validates turn and card ownership
 • Calls existing game.PlayCard() method
 • Updates game state
 • Broadcasts updates to all room players


### Phase 3: Broadcasting System

1. Create broadcast functions for game state updates
2. Integrate with existing activeGames map
3. Handle error cases and send error messages back to clients

### Phase 4: Integration Points

1. Connect to existing game logic in pkg/game/game.go
2. Use existing room management with activeRooms
3. Leverage existing session management for user authentication

### Key Implementation Details:

Message Flow:

Client WebSocket → Client.readPump → MessageRouter → handleCardPlay → Game
Logic → Broadcast → All Clients

Integration with Existing Code:

• Use activeGames[roomID] to get game state
• Use game.PlayCard() for validation and logic
• Use app.hub for WebSocket broadcasting
• Use existing GameMessage struct for communication

File Locations:

• Add message router and handlers to cmd/web/handler.go
• Modify pkg/ws/client.go readPump slightly
• No new files needed - integrate with existing structure

This approach leverages your existing infrastructure while adding the
missing game-specific WebSocket functionality.