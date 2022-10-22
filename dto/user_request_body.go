package dto

import "github.com/titoyudha/go_ecommerce_api/model"

type UserRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserRequestParams struct {
	UserID string
}

type UserResponseBody struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func ParseFromEntity(data model.User) UserResponseBody {
	return UserResponseBody{
		UserID:    data.UserID,
		FirstName: data.FirstName,
		LastName:  data.FirstName,
		Email:     data.Email,
	}
}

func ParseFromEntities(data []model.User) []UserResponseBody {
	var userResponseBodies []UserResponseBody

	for _, d := range data {
		userResponseBody := UserResponseBody{
			UserID:    d.UserID,
			FirstName: d.FirstName,
			LastName:  d.LastName,
			Email:     d.Email,
		}
		userResponseBodies = append(userResponseBodies, userResponseBody)
	}
	return userResponseBodies
}
