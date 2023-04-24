package service

import (
	"main/db"
	"main/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get Hotels
func GetHotels(c *gin.Context) {
	ctx := c.Request.Context()

	collection, err := db.Collection(db.HotelsCollection)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.Error{})
		return
	}
	defer result.Close(ctx)

	hotels := []*model.Hotel{}
	if err := result.All(ctx, &hotels); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	// Response
	c.JSON(http.StatusOK, model.Response{Data: hotels})
}

// Get hotel by id
func GetHotelById(c *gin.Context) {
	ctx := c.Request.Context()
	hotelId := c.Param("hotelId")

	hotelObjectId, err := primitive.ObjectIDFromHex(hotelId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.Error{})
		return
	}

	collection, err := db.Collection(db.HotelsCollection)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	filter := bson.M{"_id": hotelObjectId}
	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		c.JSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	hotel := &model.Hotel{}
	if err := result.Decode(hotel); err != nil {
		c.JSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	c.JSON(http.StatusOK, &model.Response{Data: hotel})
}

// Create Hotel
func CreateHotel(c *gin.Context) {
	ctx := c.Request.Context()

	now := time.Now()
	hotel := &model.Hotel{
		Statistics:        &model.Statistics{ViewCount: 0, LikeCount: 0},
		PhotoAlbums:       []string{},
		AvailableRoomDays: []string{},
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := c.ShouldBind(&hotel); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: &model.Error{}})
		return
	}

	collection, err := db.Collection(db.HotelsCollection)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	result, err := collection.InsertOne(ctx, hotel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	// Response
	resultId := result.InsertedID
	c.JSON(http.StatusOK, model.Response{Data: resultId})
}

// Update hotel by id
func UpdateHotelById(c *gin.Context) {
	ctx := c.Request.Context()
	hotelId := c.Param("hotelId")

	hotel := &model.Hotel{}
	if err := c.ShouldBind(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, &model.Error{})
		return
	}

	collection, err := db.Collection(db.HotelsCollection)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	hotelObjectId, err := primitive.ObjectIDFromHex(hotelId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.Error{})
		return
	}

	filter := bson.M{"_id": hotelObjectId}
	update := bson.M{"$set": hotel, "$currentDate": bson.M{"updatedAt": true}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := collection.FindOneAndUpdate(ctx, filter, update, options)
	if result.Err() != nil {
		c.JSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	if err := result.Decode(hotel); err != nil {
		c.JSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	// Response
	c.JSON(http.StatusOK, &model.Response{Data: hotel})
}

// Delete hotel by id
func DeleteHotelById(c *gin.Context) {
	ctx := c.Request.Context()
	hotelId := c.Param("hotelId")

	collection, err := db.Collection(db.HotelsCollection)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	hotelObjectId, err := primitive.ObjectIDFromHex(hotelId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.Error{})
		return
	}

	filter := bson.M{"_id": hotelObjectId}
	result := collection.FindOneAndDelete(ctx, filter)
	if result.Err() != nil {
		c.JSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	// Decode the deleted hotel from the result
	hotel := &model.Hotel{}
	if err := result.Decode(hotel); err != nil {
		c.JSON(http.StatusInternalServerError, &model.Error{})
		return
	}

	// Response
	c.JSON(http.StatusOK, &model.Response{Data: hotel})
}
