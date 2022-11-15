package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"log"
	"net/http"
	"time"
)

type GinServer struct {
	server *http.Server
}

func (gs *GinServer) initServer() {
	router := gin.Default()

	if config.GetString("env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	initRouter(router)

	fmt.Println(">>> Gin starting server on:", config.GetString("httpserver.default.address"))
	gs.server = &http.Server{
		Addr:    config.GetString("httpserver.default.address"),
		Handler: router,
	}
}

func (gs *GinServer) registerRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (gs *GinServer) Start() error {
	go func() {
		if err := gs.server.ListenAndServe(); err != nil {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	return nil
}

func (gs *GinServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := gs.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

func NewGinServer() *GinServer {
	s := new(GinServer)
	s.initServer()
	return s
}
