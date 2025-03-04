package orderservice

import (
	"fmt"
	"geko/internal/server"
	"net/http"
	"sync"
	"time"

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
	s.route.GET("/long", s.LongRequest)
	s.route.GET("/fast", s.FastRequest)
}

func (s *OrderService) TestOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Order test": s.serverContext.Config})
}

// slow request
func (s *OrderService) LongRequest(ctx *gin.Context) {
	t := make(chan int, 1)
	t <- 0
	wg := sync.WaitGroup{}
	fmt.Println("Request start processing")
	for {
		fmt.Println(".......")
		wg.Add(1)
		tr := <-t
		tr += 1
		fmt.Println("channel value", tr)
		if tr == 60 {
			break
		}
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			t <- tr
		}()

	}

	defer close(t)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Process done!",
	})
}

func (s *OrderService) FastRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Process done!",
	})
}
