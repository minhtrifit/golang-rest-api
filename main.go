package main

import (
	"rest_api_golang/configs"
	"rest_api_golang/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/docs/
// https://blog.logrocket.com/integrating-mongodb-go-applications/
// https://github.com/gin-contrib/cors

func main() {
	mongoClient := configs.ConnectDatabase();

	router := gin.Default();

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", controllers.HandleRunServer);

	router.GET("/albums/:id", func(ctx *gin.Context) {
		controllers.GetAlbumByID(ctx, mongoClient);
	});

	router.GET("/albums", func(ctx *gin.Context) {
		controllers.GetAlbums(ctx, mongoClient);
	});

	router.POST("/albums/add", func(ctx *gin.Context) {
		controllers.AddNewAlbum(ctx, mongoClient);
	});

	router.DELETE("/albums/delete", func(ctx *gin.Context) {
		controllers.DeleteAlbumById(ctx, mongoClient);
	});

	router.PUT("/albums/edit", func(ctx *gin.Context) {
		controllers.EditAlbum(ctx, mongoClient);
	});

	router.Run("localhost:5000");
}