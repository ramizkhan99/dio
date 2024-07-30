package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Username     string             `json:"username" bson:"username" binding:"required"`
	EmailId      string             `json:"emailId" bson:"emailId"`
	FirstName    string             `json:"firstName" bson:"firstName" binding:"required"`
	LastName     string             `json:"lastName" bson:"lastName" binding:"required"`
	CreatedAt    primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt"`
	LastModified primitive.DateTime `json:"lastModified,omitempty" bson:"lastModified"`
}
