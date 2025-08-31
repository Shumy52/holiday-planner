// Package server
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shumy52/holiday-planner/backend/internal/config"
	"github.com/shumy52/holiday-planner/backend/internal/db"
	"github.com/shumy52/holiday-planner/backend/internal/handlers"
	"github.com/shumy52/holiday-planner/backend/internal/middleware"
)

func main() {
	cfg := config.FromEnv()
	dbase, err := db.Connect(cfg.DB)
	if err != nil {
		panic(err)
	}
	auth, err := middleware.NewAuth(cfg.Issuer, cfg.Audience)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })

	api := r.Group("/api/v1")
	api.Use(auth.Require())
	vac := &handlers.VacHandler{DB: dbase}
	api.GET("/vacations/mine", vac.ListMine)
	api.POST("/vacations", vac.Create)

	_ = r.Run("0.0.0.0:" + cfg.Port)
}
