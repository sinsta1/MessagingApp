package api

import (
	"encoding/json"
	"messagingApp/backend/pkg/db"
	"messagingApp/backend/pkg/models"
	"messagingApp/backend/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateMessageHandler handles the POST /api/messages API
func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var newMessage models.Message

	// Decode the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the incoming message fields
	if newMessage.CustomerName == "" || newMessage.MessageText == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Customer name and message text are required")
		return
	}

	// Set the default status for the message
	newMessage.Status = "pending"

	// Save the message to the database
	messageID, err := db.CreateMessage(&newMessage)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create message")
		return
	}

	// Send a success response
	response := map[string]interface{}{
		"message":    "Message submitted successfully",
		"message_id": messageID,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// GetPendingMessagesHandler handles the GET /api/messages API
func GetPendingMessagesHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all pending messages from the database
	messages, err := db.GetPendingMessages()
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve messages")
		return
	}

	// Send the list of messages as a JSON response
	utils.SendJSONResponse(w, http.StatusOK, messages)
}

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the message ID from the URL
	idStr := r.URL.Path[len("/api/messages/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid message ID")
		return
	}

	// Fetch the message and responses from the database
	message, responses, err := db.GetMessageByID(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve message")
		return
	}
	if message == nil {
		utils.SendErrorResponse(w, http.StatusNotFound, "Message not found")
		return
	}

	// Create a response structure to return
	response := map[string]interface{}{
		"message":   message,
		"responses": responses,
	}

	// Send the response as JSON
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// RespondToMessageHandler handles the POST /api/messages/{id}/respond API
func RespondToMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the message ID from the URL
	vars := mux.Vars(r)
	messageID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid message ID")
		return
	}

	// Decode the incoming JSON request
	var responseData struct {
		AgentName string `json:"agent_name"`
		Response  string `json:"response"`
	}

	if err := json.NewDecoder(r.Body).Decode(&responseData); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the input fields
	if responseData.AgentName == "" || responseData.Response == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Agent name and response are required")
		return
	}

	// Save the response in the database
	if err := db.SaveResponse(messageID, responseData.AgentName, responseData.Response); err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to save response")
		return
	}

	// Update the message status to "responded"
	if err := db.UpdateMessageStatus(messageID, "responded"); err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to update message status")
		return
	}

	// Send a success response
	utils.SendJSONResponse(w, http.StatusOK, map[string]string{
		"message": "Response submitted successfully",
	})
}
