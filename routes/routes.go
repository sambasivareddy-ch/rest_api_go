package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/rest_api_go/middleware"
)

// Base function to register all the routes
func RegisterRoutes(httpServer *gin.Engine) {
	// ------------------- EVENT Routes ----------------------------
	// GET Events
	httpServer.GET("/events", getAllEventsRoute)

	// GET the event details based on EventID
	httpServer.GET("/events/:eventID", getEventByIDRoute)

	// Groups all the event authorization required routes
	eventGroupHandler := httpServer.Group("/")

	// Authorize the token before giving access to the user
	eventGroupHandler.Use(middleware.AuthorizeUser)

	// Create Events handler
	eventGroupHandler.POST("/events", createEventRoute)

	// Updates an event based on EventID passed
	eventGroupHandler.PUT("/events/:eventID", updateEventByIdRoute)

	// Deletes an event based on EventID
	eventGroupHandler.DELETE("/events/:eventID", deleteEventByIdRoute)

	// ------------------- Registration Routes ----------------------
	registerGroupHandler := httpServer.Group("/")

	// Authorize user before doing registration events
	registerGroupHandler.Use(middleware.AuthorizeUser)

	registerGroupHandler.POST("/register", registerTheUserHandler)

	registerGroupHandler.DELETE("/register", unregisterUserHandler)

	// -------------------- USER Routes -----------------------------
	// Signup Handler
	httpServer.POST("/signup", createUserRoute)

	// Login Handler
	httpServer.GET("/login", userLoginHandler)
}
