package resolver

import (
	"context"
	"graphqlGO/graph/model"
	"graphqlGO/models"
	"graphqlGO/repository"
	"strconv"
)

func CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	user := &models.User{
		Name:  name,
		Email: email,
	}

	_, err := repository.CreateUser(*user)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    strconv.Itoa(int(user.ID)),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func DeleteUser(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	duser, err := repository.DeleteUser(uint(userID))
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    strconv.Itoa(int(duser.ID)),
		Name:  duser.Name,
		Email: duser.Email,
	}, nil
}
