package repo

import (
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

// type Userdetail struct {
// 	ID          int64     `gorm:"primaryKey"`
// 	Username    string    `gorm:"column:username;unique;not null"`
// 	Password    string    `gorm:"column:password;not null"`
// 	Address     string    `gorm:"column:address;not null"`
// 	City        string    `gorm:"column:city;not null"`
// 	Pincode     int64     `gorm:"column:pincode;not null"`
// 	Phonenumber int64     `gorm:"column:phone_number; not null"`
// 	Mail        string    `gorm:"column:mail;not null"`
// 	Status      bool      `gorm:"column:status;default:true;not null"` // Boolean field, default true, to set user active or not
// 	CreatedAt   time.Time `gorm:"column:created_at;autoUpdateTime"`
// 	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
// 	UpdatedBy   *int64    `gorm:"column:deleted_by"`
// }

func (r *UserRepoImpl) SaveUserDetails(args *dto.SaveUserDetailRequest) (int64, error) {
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
