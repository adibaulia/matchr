package usecase

import (
	"fmt"
	"time"

	"github.com/adibaulia/matchr/domain"
	"github.com/adibaulia/matchr/domain/generated"
	"github.com/dgrijalva/jwt-go"
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
		Gender:          domain.MALE,
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

	ok := CheckPasswordHash(password, user.Password)
	if !ok {
		return "", fmt.Errorf("invalid password")
	}
	token, err := createToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func createToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secretKey"), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}
	customerXID, ok := claims["userID"].(string)
	if !ok {
		return "", fmt.Errorf("invalid userID claim")
	}
	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return "", fmt.Errorf("token has expired")
	}
	return customerXID, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
