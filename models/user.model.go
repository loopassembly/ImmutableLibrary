package models

import (
	"mime/multipart"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                 string    `gorm:"type:string;primary_key"`
	Name               string    `gorm:"type:varchar(100);not null"`
	Email              string    `gorm:"type:varchar(100);unique;not null"`
	Password           string    `gorm:"type:varchar(100);not null"`
	Role               *string   `gorm:"type:varchar(50);default:'user';not null"`
	Provider           *string   `gorm:"type:varchar(50);default:'local';not null"`
	Photo              *string   `gorm:"not null;default:'default.png'"`
	Verified           *bool     `gorm:"not null;default:false"`
	VerificationCode   string    `gorm:"type:varchar(100);"`
	PasswordResetToken string    `gorm:"type:varchar(100);"`
	PasswordResetAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreatedAt          time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	u.ID = uuid.String()
	return nil
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Photo           string `json:"photo"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserResponse struct
type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FilterUserRecord filters user information for response
func FilterUserRecord(user *User) UserResponse {
	id := user.ID
	return UserResponse{
		ID:        uuid.MustParse(id),
		Name:      user.Name,
		Email:     user.Email,
		Role:      *user.Role,
		Photo:     *user.Photo,
		Provider:  *user.Provider,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type Book struct {
	ID            string    `gorm:"type:string;primary_key"`
	Title         string    `gorm:"type:varchar(100);not null"`
	Author        string    `gorm:"type:varchar(100);not null"`
	Description   string    `gorm:"type:text;not null"`
	BookThumbnail string    `gorm:"type:string;not null"`
	Bookcontent   string    `gorm:"type:string;not null"`
	Genre         string    `gorm:"type:string;not null"`
	UserID        string    `gorm:"type:string;not null"`
	CreatedAt     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	b.ID = uuid.String()
	return nil
}

type BookInput struct {
	Title         string                `json:"title" validate:"required"`
	Author        string                `json:"author" validate:"required"`
	Description   string                `json:"discription" validate:"required"`
	BookThumbnail *multipart.FileHeader `json:"bookthumbnail" validate:"required"`
	Bookcontent   *multipart.FileHeader `json:"bookcontent" validate:"required"`
	Genre         string                `json:"genre" validate:"required"`
}
