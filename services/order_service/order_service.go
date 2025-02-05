package orderservice

import (
	"geko/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderService struct {
	serverContext *server.HttpServerContext
	route         *gin.RouterGroup
}

// Order constructor
func (s *OrderService) Mount(serverContext *server.HttpServerContext, route *gin.RouterGroup) {
	s.serverContext = serverContext
	s.route = route
}

// Service routes
func (s *OrderService) Routes() {
	s.route.GET("/", s.TestOrder)
}

func (s *OrderService) TestOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Order test": s.serverContext.Config})
}
