package controllers

import (
	"context"
	"rest_api_golang/configs"
	"rest_api_golang/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func GetAlbums(c *gin.Context) {
	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://admin:123@localhost:27017/?authSource=admin"));

	// col := client.Database("golang-albums").Collection("albums");
	col := models.AlbumsCollection(*client);

	// Database query
	cursor, err := col.Find(context.TODO(), bson.D{});

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

		println(configs.Purple, "Id:", id, configs.Reset);
		println(configs.Red, "Title:", album.Title, configs.Reset);
		println(configs.Blue, "Artist:", album.Artist, configs.Reset);
		println(configs.Green, "Price:", price, configs.Reset);
	}

	// Result response
	c.JSON(200, gin.H{
		"status": 200,
		"data": results,
	})
}

func GetAlbumByID(c *gin.Context) {
	client := configs.ConnectDatabase();
	coll := models.AlbumsCollection(*client);

	id := c.Query("id");
	intId, err := strconv.Atoi(id);

	println(configs.Red, "QUERY ID", id, configs.Reset);

	filter := bson.D{{Key: "id", Value: intId}};

	var result models.Album;
	err = coll.FindOne(context.TODO(), filter).Decode(&result);

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(404, gin.H{
				"status": 404,
				"message": "Album not found",
			});
			return
		}
	}else {
		c.JSON(200, gin.H{
			"status": 200,
			"message": "Find album successfully",
			"data": result,
		});
	}
}

func AddNewAlbum(c *gin.Context) {
	client := configs.ConnectDatabase();
	coll := models.AlbumsCollection(*client);

	newAlbum := models.Album{};

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return;
    }

	myAlbum := models.Album{};

	myAlbum.ID = newAlbum.ID;
	myAlbum.Title = newAlbum.Title;
	myAlbum.Artist = newAlbum.Artist;
	myAlbum.Price = newAlbum.Price;

	// Validate request body
	validate = validator.New();
	albumValid := &myAlbum;
	err := validate.Struct(albumValid);

	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"message": "Wrong format",
		});	
	}else {
		opts := options.FindOne().SetSort(bson.M{"$natural": -1})
		var lastAlbum models.Album;
		var lastAlbumBson models.AlbumBson;

		if err = coll.FindOne(context.TODO(), bson.M{}, opts).Decode(&lastAlbumBson); err != nil {
			println(configs.Red, err, configs.Reset);
		}
		
		// convert bson to struct
		bsonBytes, _ := bson.Marshal(lastAlbumBson);
		bson.Unmarshal(bsonBytes, &lastAlbumBson);

		lastAlbum.ID = lastAlbumBson.ID + 1;
		myAlbum.ID = lastAlbum.ID; // Add new album Id
		println(configs.Cyan, "New album ID:", myAlbum.ID, configs.Reset);


		// Database query
		// Add the new album to the slice.
		_, err := coll.InsertOne(context.TODO(), &myAlbum);

		// check for errors in the insertion
		if err != nil {
			c.JSON(400, gin.H{
				"status": 400,
				"message": "Insert album failed",
			})
		}
		
		// Result response
		c.JSON(201, gin.H{
			"status": 201,
			"message": "Insert album successfully",
			"data": myAlbum,
		});
	}
}