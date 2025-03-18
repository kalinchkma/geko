package geko

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	// This plug function will act as a service
	// Constructor, like service should have to implemnet this method
	// Befor plug in to server
	Mount(*HttpServerContext, *gin.RouterGroup)

	// Attach method where server route will mount
	RouteHandler()
}
