package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-server/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	gin.SetMode(gin.DebugMode)

	engine := gin.New()
	engine.Use(gin.Recovery())

	conf := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Access-Control-Allow-Origin", "Content-Type", "Accept-Encoding", "origin", "accept", "X-Requested-With", " X-CSRF-Token", "Cache-Control", "x-user-auth-token"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type"},
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{"*"},
	}

	engine.Use(cors.New(conf))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	router.Init(engine)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//if err := srv.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-quit
}
