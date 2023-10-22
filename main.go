package main

import (
	"rest_api_golang/configs"
	"rest_api_golang/controllers"

	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/docs/
// https://blog.logrocket.com/integrating-mongodb-go-applications/

func main() {
	mongoClient := configs.ConnectDatabase();

	router := gin.Default();

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

	router.POST("/albums/delete", func(ctx *gin.Context) {
		controllers.DeleteAlbumById(ctx, mongoClient);
	})

	router.Run("localhost:5000");
}