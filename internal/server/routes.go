package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.Helloworld)

	return r
}

func (s *Server) Helloworld(c *gin.Context) {
	response := make(map[string]string)

	response["message"] = "Hello world"

	c.JSON(http.StatusOK, response)
}
