package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/momomo0206/go-pizza-tracker/internal/models"
)

func main() {
	cfg := loadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cfg.DBPath)
	if err != nil {
		slog.Error("Failed to initialized database", "error", err)
		os.Exit(1)
	}

	slog.Info("Database initialized successfully")

	RegisterCustomValiidators()

	h := NewHandler(dbModel)

	router := gin.Default()

	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	setupRoutes(router, h)

	slog.Info("Server stating", "url", "http://localhost"+cfg.Port)

	router.Run(":" + cfg.Port)
}
