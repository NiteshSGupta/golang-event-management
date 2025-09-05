package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// It’s like saying:
// "Hey, I have a method called routes, it belongs to application, and when you call it, I’ll give you an HTTP handler (so you can start a server with it)."

// receiver              response in http
func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.POST("/events", app.createEvent)
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEvent)
		v1.PUT("/events/:id", app.updateEvent)
		v1.DELETE("/events/:id", app.deleteEvent)
	}

	v1.POST("/auth/register", app.registerUser)

	return g

}
