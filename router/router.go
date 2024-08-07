package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gotoeasy/glang/cmn"
	"github.com/shenzh1990/TopList/middleware/cors"
	"github.com/shenzh1990/TopList/pkg/settings"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	cmn.Info("Router Init")
	r := initRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        r,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			cmn.Info(fmt.Sprintf("Listen: %s\n", err))
		}
	}()
	cmn.Info(fmt.Sprintf("Listen: http://127.0.0.1:%d\n", settings.HTTPPort))
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	cmn.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		cmn.Error("Server Shutdown:", err)
	}
	cmn.Info("Server exiting")
}
func initRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(cors.Cors())
	gin.SetMode(settings.RunMode)

	//auth := r.Group("/api")
	//{
	//	auth.POST("/login", api.Login)
	//	auth.POST("/register", api.Register)
	//	auth.POST("/attach/upload", api.FileUpload)
	//	auth.GET("/chat", api.Chat)
	//}
	//auth.Use(jwtutil.AuthorizedMiddelware(settings.JwtSecret))
	//{
	//	auth.POST("/contact/addfriend", api.AddFriend)
	//	auth.GET("/contact/listfriends", api.ListFriends)
	//
	//}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	return r
}
