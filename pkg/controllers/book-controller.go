package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/navindu-sachintha/go-bookstore/internal/database"
	"github.com/navindu-sachintha/go-bookstore/pkg/config"
	"github.com/navindu-sachintha/go-bookstore/pkg/utils"
)

type Handler struct {
	Queries *database.Queries
}

func NewHandler() *Handler {
	db := config.GetDB()
	return &Handler{
		Queries: database.New(db),
	}
}

func (h *Handler) HandleCreateBook(c *gin.Context) {
	var req struct {
		Isbn        string `json:"isbn"`
		Name        string `json:"name"`
		Author      string `json:"author"`
		Publication string `json:"publication"`
	}

	utils.ParseBody(c.Request, &req)

	book, err := h.Queries.CreateBook(context.Background(), database.CreateBookParams{
		ID:          uuid.New(),
		Isbn:        req.Isbn,
		Name:        req.Name,
		Author:      req.Author,
		Publication: req.Publication,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": book})
}

func (h *Handler) HandleListBooks(c *gin.Context) {
	books, err := h.Queries.ListBooks(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)

}

func (h *Handler) HandleGetBook(c *gin.Context) {

	id := c.Param("id")
	uuID, _ := uuid.Parse(id)

	book, err := h.Queries.GetBook(context.Background(), uuID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusFound, book)
}

func (h *Handler) HanldeUpdateBook(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	id := c.Param("id")
	uuID, _ := uuid.Parse(id)

	utils.ParseBody(c.Request, &req)
	err := h.Queries.UpdateBookName(context.Background(), database.UpdateBookNameParams{
		ID:        uuID,
		Name:      req.Name,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Bookname updated"})
}

func (h *Handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	uuID, _ := uuid.Parse(id)

	err := h.Queries.DeleteBook(context.Background(), uuID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"OK": "book deleted"})

}
