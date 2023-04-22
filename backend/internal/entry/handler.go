package entry

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *EntryService
}

func NewHandler(entryService *EntryService) *Handler {
	return &Handler{
		Service: entryService,
	}
}

func (h *Handler) GetAllEntries(c *gin.Context) {
	entries, err := h.Service.GetAllEntries()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}

func (h *Handler) GetEntryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid entry ID"})
		return
	}

	entry, err := h.Service.GetEntryByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "entry not found"})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func (h *Handler) CreateEntry(c *gin.Context) {
	var req *CreateEntryRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	entry, err := h.Service.CreateEntry(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, entry)
}

func (h *Handler) UpdateEntry(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid entry ID"})
	// 	return
	// }

	var req *UpdateEntryRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	entry, err := h.Service.UpdateEntry(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "entry not found"})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func (h *Handler) DeleteEntry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid entry ID"})
		return
	}

	if err := h.Service.DeleteEntry(id); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "entry not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
