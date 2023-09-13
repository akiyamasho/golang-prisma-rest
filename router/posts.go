package router

import (
	"demo/handler"
	"demo/prisma/db"

	"github.com/labstack/echo/v4"
)

func PostsRouter(e *echo.Echo, dbClient *db.PrismaClient) {
	postHandler := handler.NewPostsHandler(dbClient)
	g := e.Group("/posts")
	g.GET("", postHandler.ShowAll)
	g.GET(":id", postHandler.Show)
	g.POST("", postHandler.Create)
}
