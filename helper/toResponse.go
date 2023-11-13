package helper

import (
	"api-doorsec/model/response"
	"api-doorsec/model/structure"
)

//func ToCategoryResponse(category domain.Category) api.CategoryResponse {
//	return api.CategoryResponse{
//		Id:   category.Id,
//		Name: category.Name,
//	}
//}
//
//func ToCategoryResponses(categories []domain.Category) []api.CategoryResponse {
//	var categoryResponses []api.CategoryResponse
//
//	for _, category := range categories {
//		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
//	}
//
//	return categoryResponses
//}

func ToUsersResponse(user structure.Users) response.Users {
	return response.Users{
		Id:        user.Id,
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
