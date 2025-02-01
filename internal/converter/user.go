package converter

import (
	"user-api/internal/dto"
	"user-api/internal/model"
	"user-api/pkg/gensqlc"
	"user-api/pkg/proto/gengrpc"
)

func ModelUserToDto(user model.User) dto.ResponseUserDto {
	return dto.ResponseUserDto{
		ID:            user.ID,
		Email:         user.Email,
		FirstName:     user.FirstName,
		MiddleName:    user.MiddleName,
		LastName:      user.LastName,
		Photo:         user.Photo,
		Phone:         user.Phone,
		Role:          user.Role,
		Status:        user.Status,
		Address:       user.Address,
		VerifiedToken: user.VerifiedToken,
		LastActivity:  user.LastActivity,
		IsVerified:    user.IsVerified,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

func SqlcUserToModel(user *gensqlc.User) model.User {
	return model.User{
		ID:            user.ID,
		Email:         user.Email,
		FirstName:     user.FirstName,
		MiddleName:    user.MiddleName,
		LastName:      user.LastName,
		Photo:         user.Photo,
		PasswordHash:  user.PasswordHash,
		Phone:         user.Phone,
		Role:          user.Role,
		Status:        user.Status,
		Address:       user.Address,
		VerifiedToken: user.VerifiedToken,
		LastActivity:  user.LastActivity,
		IsVerified:    user.IsVerified,
		CreatedAt:     user.CreatedAt.String(),
		UpdatedAt:     user.UpdatedAt.String(),
	}
}

func DtoUserToGrpc(user dto.ResponseUserDto) *gengrpc.User {
	return &gengrpc.User{
		Id:            user.ID.String(),
		Email:         user.Email,
		FirstName:     user.FirstName,
		MiddleName:    user.MiddleName,
		LastName:      user.LastName,
		Photo:         user.Photo,
		Phone:         user.Phone,
		Role:          user.Role,
		Status:        user.Status,
		Address:       user.Address,
		VerifiedToken: user.VerifiedToken,
		LastActivity:  user.LastActivity,
		IsVerified:    user.IsVerified,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
