package controllers

import (
	"context"
	"rest_api_golang/configs"
	"rest_api_golang/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAlbums(c *gin.Context) {
	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://admin:123@localhost:27017/?authSource=admin"));

	// col := client.Database("golang-albums").Collection("albums");
	col := models.AlbumsCollection(*client);

	cursor, err := col.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		panic(err);
	}

	// convert the cursor result to bson
	results := []bson.M{};
	// check for errors in the conversion
	if err := cursor.All(context.TODO(), &results); err != nil {
		println(configs.Red, err, configs.Reset);
	}
	
	var Albums = []models.Album{};

	for _, result := range results {
		var album models.Album;
		var albumBson models.AlbumBson;

		// convert bson to struct
		bsonBytes, _ := bson.Marshal(result);
		bson.Unmarshal(bsonBytes, &albumBson);

		album.ID = albumBson.ID;
		album.Title = albumBson.Title;
		album.Artist = albumBson.Artist;
		album.Price = albumBson.Price;

		// Insert to slice
		Albums = append(Albums, album);
	}
	
	for _, album := range Albums {
		id := strconv.Itoa(album.ID);
		price := strconv.FormatFloat(album.Price, 'f', -1, 64);

		println(configs.Purple, id, configs.Reset);
		println(configs.Red, album.Title, configs.Reset);
		println(configs.Blue, album.Artist, configs.Reset);
		println(configs.Green, price, configs.Reset);
	}

	// Result response
	c.JSON(200, gin.H{
		"status": 200,
		"data": results,
	})
}

func GetAlbumByID(c *gin.Context) {

}

func AddNewAlbum(c *gin.Context) {
	client := configs.ConnectDatabase();
	coll := models.AlbumsCollection(*client);

	var newAlbum models.Album;

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return;
    }

	myAlbum := models.Album{};

	myAlbum.ID = newAlbum.ID;
	myAlbum.Title = newAlbum.Title;
	myAlbum.Artist = newAlbum.Artist;
	myAlbum.Price = newAlbum.Price;

    // Add the new album to the slice.
	_, _ = coll.InsertOne(context.TODO(), &myAlbum);
    
	// Result response
	c.JSON(200, gin.H{
		"status": 200,
		"message": "Insert album successfully",
		"data": myAlbum,
	})
}