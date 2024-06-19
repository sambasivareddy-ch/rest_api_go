package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/rest_api_go/models"
)

func getAllEventsRoute(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to fetch the data from events table",
		})
	}

	ctx.JSON(http.StatusOK, events)
}

func createEventRoute(ctx *gin.Context) {
	var event models.Event

	// Internally Gin look into body & store the data in the event variable
	// So basically, request must have same format as Event struct, if any fields Gin will set to Default Value
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse the request data",
		})
		return // Return to prevent from further processing
	}

	event.UserID = ctx.GetInt64("userId")

	if err1 := event.Save(); err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to insert the data into events table",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":   "Event created!!",
		"eventname": event.EventName,
	})
}

func getEventByIDRoute(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("eventID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to process eventId",
		})
		return
	}

	targetEvent, err1 := models.GetEventByID(eventId)

	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("unable to get the event details with id: %d", eventId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event Details fetched successfully",
		"event":   targetEvent,
	})
}

func updateEventByIdRoute(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("eventID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to process eventId",
		})
		return
	}

	targetEvent, err1 := models.GetEventByID(eventId)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("unable to get the event details with id: %d", eventId),
		})
		return
	}

	// Compare UserId from fetched event & authorized user if not matched return from here without processing
	authorizedUserId := ctx.GetInt64("userId")
	if targetEvent.UserID != authorizedUserId {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("permission denied to update the event details with id: %d", eventId),
		})
		return
	}

	var updatedEvent models.Event
	if err = ctx.ShouldBindJSON(&updatedEvent); err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse the request data",
		})
		return // Return to prevent from further processing
	}

	if err = updatedEvent.UpdateEvent(eventId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("unable to update the event details with id: %d", eventId),
		})
		return
	}

	// Success
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event Updated",
		"event":   updatedEvent,
	})
}

func deleteEventByIdRoute(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("eventID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to process eventId",
		})
		return
	}

	tobeDeletedEvent, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("unable to get the event details with id: %d", eventId),
		})
		return
	}

	// Compare UserId from fetched event & authorized user if not matched return from here without processing
	authorizedUserId := ctx.GetInt64("userId")
	if tobeDeletedEvent.UserID != authorizedUserId {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("permission denied to delete the event details with id: %d", eventId),
		})
		return
	}

	if err = models.DeleteEventByID(eventId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("unable to delete the event details with id: %d", eventId),
		})
		return
	}

	// Success
	ctx.JSON(http.StatusCreated, gin.H{
		"message":   "Event Deleted",
		"eventName": tobeDeletedEvent.EventName,
	})
}
