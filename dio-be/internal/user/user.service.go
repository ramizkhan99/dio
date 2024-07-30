package user

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ramizkhan99/dio-be/internal/db"
	"github.com/ramizkhan99/dio-be/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	NewUser(string, string, string, string, string)
	DeleteUser(string) User
	GetUserByID(string) User
	GetUserByUsername(string) User
	UpdateUser(string, User) User
}

const COLLECTION = "users"

// TODO: Use receivers for this

func NewUser(ctx *gin.Context) (User, error) {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		logger.GetSugaredLogger().Errorf("failed to parse user data %w", err)
		return User{}, err
	}
	user.ID = primitive.NewObjectID()
	currentDatetime := primitive.NewDateTimeFromTime(time.Now())
	user.CreatedAt = currentDatetime
	user.LastModified = currentDatetime
	_, err = db.Insert[User](COLLECTION, user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUserByID(id string) (User, error) {
	user := User{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.GetSugaredLogger().Errorf("invalid object id :: %v", err)
		return user, err
	}
	collection := db.Connect().Database("dio").Collection(COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"_id": objectId})
	err = result.Decode(&user)
	if err != nil {
		logger.GetSugaredLogger().Errorf("No user found :: %+v", err.Error())
		return user, err
	}
	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	user := User{}
	collection := db.Connect().Database("dio").Collection(COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"username": username})
	err := result.Decode(&user)
	return user, err
}

func (user *User) UpdateUsername(username string) error {
	collection := db.Connect().Database("dio").Collection(COLLECTION)
	_, err := collection.UpdateByID(context.TODO(), user.ID, bson.D{{Key: "$set", Value: bson.D{{Key: "username", Value: username}} }})
	user.Username = username
	return err
}

func DeleteUser(id string) (User, error) {
	return User{}, nil
}
