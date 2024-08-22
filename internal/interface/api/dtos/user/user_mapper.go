package user

import "Or/Library/internal/domain/user"

func MapToUserDTO(user *user.User) UserDTO {
	return UserDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
	}
}

func MapToDomainUserFromDTO(dto *UserDTO) *user.User {
	return &user.User{
		ID:       dto.ID,
		Name:     dto.Name,
		Email:    dto.Email,
		Role:     dto.Role,
		Password: dto.Password,
	}
}
