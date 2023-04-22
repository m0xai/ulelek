package entry

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEntryRouter(db *sql.DB) http.Handler {
	entryRepo := NewEntryRepo(db)
	entryService := NewEntryService(entryRepo)
	entryHandler := NewHandler(entryService)

	router := gin.Default()
	router.GET("api/v1/entries", entryHandler.GetAllEntries)
	router.GET("api/v1/entries/:id", entryHandler.GetEntryByID)
	router.POST("api/v1/entries", entryHandler.CreateEntry)
	router.PUT("api/v1/entries/:id", entryHandler.UpdateEntry)
	router.DELETE("api/v1/entries/:id", entryHandler.DeleteEntry)

	return router
}
