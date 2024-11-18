package routes

import (
	"net/http"
	"strconv"

	"github.com/MacBie2k/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, &event)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.UserId = context.GetInt64("userId")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	existingEvent, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event"})
		return
	}

	userId := context.GetInt64("userId")

	if userId != existingEvent.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not update event"})
		return
	}

	event.Id = id
	event.UserId = userId
	err = event.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated!"})
}

func deleteEvent(context *gin.Context) {

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
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not delete event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})
}
