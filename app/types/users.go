package types

import (
	"github.com/google/uuid"
)

// RegisterRequestDto defines the structure for the /api/user/signup/ POST body
type RegisterRequestDto struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,min=10,max=255"`
}

// BasicProfileDto defines the structure for the /api/v1/profile/ PATCH body
type BasicProfileDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// ChangePasswordDto defines the structure for the /api/v1/profile/password/ PATCH body
type ChangePasswordDto struct {
	CurrentPassword string `json:"currentPassword" binding:"required,min=10,max=255"`
	NewPassword     string `json:"newPassword" binding:"required,min=10,max=255"`
}

// GetUserResponseDto defines the Response body for GET /admin/users/
type GetUserResponseDto struct {
	ID        uuid.UUID `json:"ID" binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email" binding:"required"`
}
