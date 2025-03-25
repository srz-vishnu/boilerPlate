package repo

import (
	"fmt"
	"pjt1/app/domain"
	"pjt1/app/dto"

	"gorm.io/gorm"
)

type UserRepo interface {
	SaveUserDetails(args *dto.SaveUserDetailRequest) (int64, error)
	GetUserByUsername(username string) (*domain.Userdetail, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

func (r *UserRepoImpl) SaveUserDetails(args *dto.SaveUserDetailRequest) (int64, error) {

	// Check if the email already exists
	var existingUser domain.Userdetail
	err := r.db.Table("userdetails").Where("mail = ?", args.Mail).First(&existingUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	if err == nil {
		// Email is already in use (record found), return an error
		return 0, fmt.Errorf("email %s is already in use", args.Mail)
	}

	user := domain.Userdetail{
		//ID:       args.UserID,
		Address:  args.Address,
		City:     args.City,
		Mail:     args.Mail,
		Username: args.UserName,
		Password: args.Password,
		Pincode:  args.Pincode,
	}
	//GORM's Create method to insert the new user
	if err := r.db.Table("userdetails").Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepoImpl) GetUserByUsername(username string) (*domain.Userdetail, error) {
	var user domain.Userdetail
	if err := r.db.Table("userdetails").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
