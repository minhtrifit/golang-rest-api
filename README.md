> ## Golang Restfull API Docker Document

## RUN DOCKER: `docker-compose -f .\docker-compose.yml up`

## RUN SERVER: `go run main.go`

## Open Mongo-express browser: `http://localhost:8081`
### Username: `root`
### Password: `root`

## API Listing

### Get All Albums:

**Request**

`GET: http://localhost:5000/albums`

**Response**

```php
{
  "data": [
    {
      "_id": "65348f16fea56e0d94903eab",
      "artist": "John Coltrane",
      "id": 1,
      "price": 56.99,
      "title": "Blue Train"
    },
    {
      "_id": "65348f2efea56e0d94903ead",
      "artist": "Gerry Mulligan",
      "id": 2,
      "price": 17.99,
      "title": "Blue Train"
    },
    {
      "_id": "65348f4cfea56e0d94903eaf",
      "artist": "Sarah Vaughan",
      "id": 3,
      "price": 39.99,
      "title": "Sarah Vaughan and Clifford Brown"
    }
  ],
  "status": 200
}
```

### Get album by id

**Request Body**

`GET: http://localhost:5000/albums/query?id=1`

**Response**

```php
{
  "data": {
    "id": 1,
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
  },
  "message": "Find album successfully",
  "status": 200
}
```

### Insert New Album:

`POST: http://localhost:5000/albums/add`

**Request Body**

```php
{
  "title": "Nơi Nào Có Tình Yêu",
  "artist": "Minh Trí",
  "price": 150.00
}
```

**Response**

```php
{
  "data": {
    "id": 4,
    "title": "Nơi Nào Có Tình Yêu",
    "artist": "Minh Trí",
    "price": 150
  },
  "message": "Insert album successfully",
  "status": 200
}
```

### Delete Album By Id:

`DELETE: http://localhost:5000/albums/delete`

**Request Body**

```php
{
  "id": 4
}
```

**Response**

```php
{
  "albums": 4,
  "message": "Delete album successfully",
  "status": 200
}
```

### Edit Album:

`PUT: http://localhost:5000/albums/edit`

**Request Body**

```php
{
  "id": 4,
  "title": "Nơi Không Có Em 2",
  "artist": "Minh Trí",
  "price": 150.69
}
```

**Response**

```php
{
  "data": {
    "MatchedCount": 1,
    "ModifiedCount": 1,
    "UpsertedCount": 0,
    "UpsertedID": null
  },
  "message": "Edit album successfully",
  "status": 200
}
```

### Showcases:

**GET Method**

![GET](./showcase/get.png "GET Method")

**POST Method**

![POST](./showcase/post.png "POST Method")