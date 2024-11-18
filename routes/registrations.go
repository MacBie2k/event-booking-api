package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event"})
		return
	}

	userId := context.GetInt64("userId")
	registration, err := models.GetRegistrationByUserAndEventId(userId, event.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register to event"})
		return
	}

	if registration != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "Already registered to this event"})
		return
	}

	registration = &models.Registration{UserId: userId, EventId: event.Id}

	err = registration.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register to event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered successfully"})

}

func cancelRegistration(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event"})
		return
	}

	userId := context.GetInt64("userId")
	registration, err := models.GetRegistrationByUserAndEventId(userId, event.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration to event"})
		return
	}

	if registration == nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Registration doesn't exist"})
		return
	}

	err = registration.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration to event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled successfully"})
}
