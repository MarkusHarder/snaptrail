package structs

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AdminRole = "admin"
	Readonly  = "readonly"
)

type Session struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Thumbnail   Thumbnail `gorm:"foreignKey:SessionID;references:ID" json:"thumbnail"`
	Name        string    `gorm:"type:text;not null" json:"sessionName"`
	Subtitle    string    `gorm:"type:text;not null" json:"subtitle"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Date        time.Time `gorm:"type:text;not null" json:"date"`
	Published   bool      `gorm:"type:bool;not null" json:"published"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Thumbnail struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	SessionID uint64 `json:"sessionId"`
	Filename  string `gorm:"type:text;not null" json:"filename"`
	MimeType  string `gorm:"type:text;not null" json:"mimeType"`
	ExifMetadata
	Data      []byte    `gorm:"type:bytea;not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ExifMetadata struct {
	CameraModel string  `gorm:"type:text;" json:"cameraModel"`
	Make        string  `gorm:"type:text;" json:"make"`
	LensModel   string  `gorm:"type:text;" json:"lensModel"`
	Exposure    string  `gorm:"type:text;" json:"exposure"`
	DateTime    string  `gorm:"type:text;" json:"dateTime"`
	Aperture    float64 `gorm:"type:double precision;" json:"aperture"`
	ISO         int     `gorm:"type:int;not null" json:"iso"`
	FocalLength float64 `gorm:"type:double precisionn;" json:"fc"`
}

type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:text;not null" json:"username"`
	Password string `gorm:"type:text;not null" json:"password"`
	Version  int64  `gorm:"type:int;not null" json:"-"`
	Role     string `gorm:"type:text;not null" json:"role"`
}

type PasswordChange struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassowrd string `json:"newPassword"`
}
type CustomClaims struct {
	Role    string `json:"role"`
	Version int64  `json:"version"`
	jwt.RegisteredClaims
}
