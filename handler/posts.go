package handler

import (
	"context"
	"demo/prisma/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostsHandler struct {
	client *db.PrismaClient
}

type PostsResponse struct {
	db.PostModel
}

func NewPostsHandler(client *db.PrismaClient) *PostsHandler {
	return &PostsHandler{client}
}

func (h *PostsHandler) ShowAll(c echo.Context) error {
	todos, err := h.client.Post.FindMany().Exec(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, "error getting Posts")
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *PostsHandler) Show(c echo.Context) error {
	todo, err := h.client.Post.FindUnique(db.Post.ID.Equals(c.Param("id"))).Exec(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *PostsHandler) Create(c echo.Context) error {
	var post db.PostModel
	if err := c.Bind(&post); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	created, err := h.client.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		nil,
	).Exec(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, created)
}
