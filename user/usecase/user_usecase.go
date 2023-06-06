package usecase

import (
	"fmt"
	"time"

	"github.com/adibaulia/matchr/domain"
	"github.com/adibaulia/matchr/domain/generated"
	"github.com/adibaulia/matchr/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUseCase {
	return &UserUsecase{repo}
}

func (u *UserUsecase) RegisterUser(user domain.User) error {
	hashedPass, _ := HashPassword(user.Password)
	userRepo := &generated.User{
		Username:           user.Username,
		Email:              user.Email,
		Password:           hashedPass,
		VerificationStatus: false,
		UserStatus:         string(domain.FREE),
	}

	userID, err := u.repo.CreateUser(*userRepo)
	if err != nil {
		return err
	}
	birthday, err := time.Parse("02-01-2006", user.Profile.DateOfBirth)
	if err != nil {
		return err
	}
	profileRepo := generated.Profile{
		Name:            user.Profile.Name,
		Gender:          user.Profile.Gender,
		DateOfBirth:     birthday,
		Bio:             user.Profile.Bio,
		ProfileImageURL: user.Profile.ProfileImageURL,
	}

	_, err = u.repo.CreateProfileUser(userID, profileRepo)
	if err != nil {
		return err
	}

	return nil

}
func (u *UserUsecase) LoginUser(username, password string) (string, error) {
	user, err := u.repo.GetUserByUserName(username)
	if err != nil {
		return "", err
	}
	if user.Username == "" {
		return "", fmt.Errorf("invalid username")
	}

	ok := CheckPasswordHash(password, user.Password)
	if !ok {
		return "", fmt.Errorf("invalid password")
	}
	token, err := token.CreateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
