package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"name,omitempty" json:"name,omitempty"`
	Description       string             `bson:"description,omitempty" json:"description,omitempty"`
	ImageUrl          string             `bson:"imageUrl,omitempty" json:"imageUrl,omitempty"`
	CoverPhotoUrl     string             `bson:"coverPhotoUrl,omitempty" json:"coverPhotoUrl,omitempty"`
	PhotoAlbums       []string           `bson:"photoAlbums,omitempty" json:"photoAlbums,omitempty"`
	PhoneNumber       string             `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	RoomPrice         int                `bson:"roomPrice,omitempty" json:"roomPrice,omitempty"`
	RoomMaxPrice      int                `bson:"roomMaxPrice,omitempty" json:"roomMaxPrice,omitempty"`
	AvailableRoomDays []string           `bson:"availableRoomDays,omitempty" json:"availableRoomDays,omitempty"`
	SocialUrls        *SocialUrls `bson:"socialUrls,omitempty" json:"socialUrls,omitempty"`
	Statistics        *Statistics `bson:"statistics,omitempty" json:"statistics,omitempty"`
	IsPublic          bool        `bson:"isPublic,omitempty" json:"isPublic,omitempty"`
	IsActive          bool        `bson:"isActive,omitempty" json:"isActive,omitempty"`
	CreatedAt         time.Time   `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt         time.Time   `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type SocialUrls struct {
	Facebook string
	Line     string
}

type Statistics struct {
	ViewCount int `bson:"viewCount" json:"viewCount" binding:"required"`
	LikeCount int `bson:"likeCount" json:"likeCount" binding:"required"`
}
