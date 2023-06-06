package postgre

import (
	"github.com/adibaulia/matchr/domain"
	"github.com/adibaulia/matchr/domain/generated"
	"gorm.io/gorm"
)

type (
	postgreUserRepo struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &postgreUserRepo{db}
}

func (r *postgreUserRepo) CreateUser(user generated.User) (string, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
func (r *postgreUserRepo) CreateProfileUser(userID string, profile generated.Profile) (string, error) {
	profile.UserID = userID

	err := r.db.Create(&profile).Error
	if err != nil {
		return "", err
	}
	return profile.ID, nil
}
func (r *postgreUserRepo) GetUserByUserName(username string) (*generated.User, error) {
	var user generated.User
	err := r.db.Where(&generated.User{Username: username}).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
