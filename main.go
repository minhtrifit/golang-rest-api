package main

import (
	"rest_api_golang/configs"
	"rest_api_golang/controllers"

	"github.com/gin-gonic/gin"
)

// https://blog.logrocket.com/integrating-mongodb-go-applications/

func main() {
	configs.ConnectDatabase();

	router := gin.Default();

	router.GET("/", controllers.HandleRunServer);
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.GET("/albums", controllers.GetAlbums);
	router.POST("/albums/add", controllers.AddNewAlbum);

	router.Run("localhost:5000");
}