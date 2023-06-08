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

func (r *postgreUserRepo) GetUserByUserID(userID string) (*generated.User, error) {
	var user generated.User
	err := r.db.Where(&generated.User{ID: userID}).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *postgreUserRepo) FindPotentialMatchr(userID string) (*domain.UserProfile, error) {
	var userProfile domain.UserProfile
	err := r.db.
		Model(&generated.User{}).
		Select("users.id, users.username, users.email, users.user_status, users.verification_status, p.name, p.date_of_birth, p.gender, p.bio, p.profile_image_url").
		Joins("JOIN profiles p ON users.id = p.user_id").
		Where("users.id NOT IN (?) AND users.id NOT IN (?) AND p.gender NOT IN (?)",
			r.db.Table("swipes").Select("swiped_id").Where("swiper_id = ?", userID),
			r.db.Table("swipes").Select("swiper_id").Where("swiped_id = ? AND swipe_direction = false", userID),
			r.db.Table("profiles").Select("gender").Where("user_id = ?", userID),
		).
		Scan(&userProfile).Error
	if err != nil {
		return nil, err
	}
	if userProfile.Username == "" {
		return nil, nil
	}

	return &userProfile, nil
}

func (r *postgreUserRepo) GetProfileByUserID(userID string) (*generated.Profile, error) {
	var profile generated.Profile

	err := r.db.Where(&generated.Profile{UserID: userID}).Find(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
