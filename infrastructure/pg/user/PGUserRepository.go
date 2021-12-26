package user

import (
	"errors"

	"gorm.io/gorm"

	u "nrsDDD/domain/models/user"
)

type User struct {
	gorm.Model
	Id   string `gorm:"primaryKey"`
	Name string
}

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*UserRepository, error) {
	ur := &UserRepository{
		db: db,
	}
	return ur, nil
}

func (ur *UserRepository) Save(user u.User) error {
	result := ur.db.Create(&User{
		Id:   user.Id.Value,
		Name: user.Name.Value,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) FindById(userId string) (*u.User, error) {
	var user User
	result := ur.db.First(&user, "id = ?", userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}
	ui, _ := u.NewUserId(user.Id)
	un, _ := u.NewUserName(user.Name)
	return &u.User{Id: *ui, Name: *un}, nil
}

func (ur *UserRepository) FindByName(userName string) (*u.User, error) {
	var user User
	result := ur.db.Where("name = ?", userName).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}
	ui, _ := u.NewUserId(user.Id)
	un, _ := u.NewUserName(user.Name)
	return &u.User{Id: *ui, Name: *un}, nil
}

func (ur *UserRepository) Update(user u.User) error {
	result := ur.db.Model(&User{}).Where("id = ?", user.Id.Value).Update("name", user.Name.Value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) Delete(user u.User) error {
	result := ur.db.Delete(&User{
		Id: user.Id.Value,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
