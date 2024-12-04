package routes

import (
	"example.com/main/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEventById)
	// server.POST("/events", middlewares.Authenticate, createEvents)
	// server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.GET("/products", getProducts)
	server.GET("/products/:id", getProductById)
	server.POST("/products", createProducts)
	server.PUT("/products/:id", updateProduct)
	server.DELETE("/products/:id", deleteProduct)

	server.POST("/signup", signup)
	server.POST("/login", login)
}