package domain

import (
	"time"

	"github.com/adibaulia/matchr/domain/generated"
)

type (
	User struct {
		ID                 string  `json:"id"`
		Username           string  `json:"username"`
		Email              string  `json:"email"`
		Password           string  `json:"password"`
		UserStatus         string  `json:"user_status"`
		VerificationStatus bool    `json:"verification_status"`
		Profile            Profile `json:"profile"`
	}

	Profile struct {
		Name            string `json:"name"`
		DateOfBirth     string `json:"date_of_birth"`
		Gender          string `json:"gender"`
		Bio             string `json:"bio"`
		ProfileImageURL string `json:"profile_image_url"`
	}

	UserProfile struct {
		ID                 string    `json:"id"`
		Username           string    `json:"username"`
		Email              string    `json:"email"`
		UserStatus         string    `json:"user_status"`
		VerificationStatus bool      `json:"verification_status"`
		Name               string    `json:"name"`
		DateOfBirth        time.Time `json:"date_of_birth"`
		Gender             string    `json:"gender"`
		Bio                string    `json:"bio"`
		ProfileImageURL    string    `json:"profile_image_url"`
	}

	UserUsecase interface {
		RegisterUser(user User) error
		LoginUser(username, password string) (string, error)
		FindPotentialMatchr(userID string) (*User, error)
	}

	UserRepository interface {
		CreateUser(user generated.User) (string, error)
		CreateProfileUser(userID string, profile generated.Profile) (string, error)
		GetUserByUserName(userName string) (*generated.User, error)
		GetUserByUserID(userName string) (*generated.User, error)
		FindPotentialMatchr(userID string) (*UserProfile, error)
		GetProfileByUserID(userID string) (*generated.Profile, error)
	}
)
