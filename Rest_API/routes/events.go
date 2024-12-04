package routes

import (
	"net/http"
	"strconv"

	"example.com/main/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	//Get all events
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	//Create new event
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse Json data"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = int(userId)
	
	err = event.AddEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create events. Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Create successful", "event": event})
}

func getEventById(context *gin.Context) {
	//Get id param
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse event id"})
		return
	}

	//Get event by id
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	//Get id param
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse event id"})
		return
	}

	//Get event by id
	event, err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch event"})
		return
	}

	//Check authorize to update
	userId := context.GetInt64("userId")
	if event.UserID != int(userId){
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize to update event"})
		return
	}

	//Update event
	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse request data"})
		return
	}
	updateEvent.ID = eventId
	err = updateEvent.UpdateEvent()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't update event"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Update successfully"})
}

func deleteEvent(context *gin.Context){
	//Get id param
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse event id"})
		return
	}

	//Get event by id
	event, err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch event"})
		return
	}

	//Check authorize to update
	userId := context.GetInt64("userId")
	if event.UserID != int(userId){
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize to delete event"})
		return
	}

	//Delete event
	err = event.DeleteEvent()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete succesfully"})
}
