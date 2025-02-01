package converter

import (
	"user-api/internal/dto"
	"user-api/pkg/proto/gengrpc"
)

func UserResponseDtoToMeResponseDto(req dto.ResponseUserDto) dto.MeResponseDto {
	return dto.MeResponseDto{
		ID:           req.ID,
		Email:        req.Email,
		FirstName:    req.FirstName,
		MiddleName:   req.MiddleName,
		LastName:     req.LastName,
		Role:         req.Role,
		Status:       req.Status,
		Phone:        req.Phone,
		Address:      req.Address,
		LastActivity: req.LastActivity,
		IsVerified:   req.IsVerified,
		Photo:        req.Photo,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}
}

func MeResponseDtoToAuthUser(dto dto.MeResponseDto) *gengrpc.AuthUser {
	return &gengrpc.AuthUser{
		Id:           dto.ID.String(),
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		MiddleName:   dto.MiddleName,
		Email:        dto.Email,
		Role:         dto.Role,
		Status:       dto.Status,
		Phone:        dto.Phone,
		Address:      dto.Address,
		LastActivity: dto.LastActivity,
		IsVerified:   dto.IsVerified,
		Photo:        dto.Photo,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
	}
}
