package repository

import (
	"context"
	"errors"
	"graphqlGO/config"
	"graphqlGO/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUser() ([]models.User, error) {

	var users []models.User

	collection := config.DB.Collection("users")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		var user models.User

		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func CreateUser(user models.User) (models.User, error) {

	collection := config.DB.Collection("users")

	// var lastUser models.User

	// opts := options.FindOne().
	// 	SetSort(bson.D{{Key: "id", Value: -1}})

	// err := collection.FindOne(
	// 	context.Background(),
	// 	bson.M{},
	// 	opts,
	// ).Decode(&lastUser)

	// // First user
	// if err != nil {
	// 	user.ID = 1
	// } else {
	// 	user.ID = lastUser.ID + 1
	// }

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func DeleteUser(id uint) (*models.User, error) {

	collection := config.DB.Collection("users")
	var user models.User

	err := collection.FindOneAndDelete(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	_, err = collection.DeleteOne(
		context.Background(),
		bson.M{"id": id},
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
