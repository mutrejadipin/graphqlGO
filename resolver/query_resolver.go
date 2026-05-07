package resolver

import (
	"context"
	"graphqlGO/graph/model"
	"graphqlGO/repository"
	"strconv"
)

func User(ctx context.Context) ([]*model.User, error) {
	users, err := repository.GetUser()
	if err != nil {
		return nil, err
	}
	var result []*model.User
	for _, user := range users {
		result = append(result, &model.User{
			ID:    strconv.Itoa(int(user.ID)),
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return result, nil
}
