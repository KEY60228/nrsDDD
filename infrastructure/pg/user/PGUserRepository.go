package user

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	u "nrsDDD/domain/models/user"
)

type User struct {
	gorm.Model
	Id   string `gorm:"primaryKey"`
	Name string
}

type UserRepository struct {
	dsn string
}

func New() (*UserRepository, error) {
	dsn := fmt.Sprintf("host=pgsql dbname=nrsDDD user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	ur := &UserRepository{
		dsn: dsn,
	}
	return ur, nil
}

func (ur *UserRepository) Save(user u.User) error {
	db, err := gorm.Open(postgres.Open(ur.dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	result := db.Create(&User{
		Id:   user.Id.Value,
		Name: user.Name.Value,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) FindById(userId string) (*u.User, error) {
	db, err := gorm.Open(postgres.Open(ur.dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	var user User
	result := db.First(&user, "id = ?", userId)
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
	db, err := gorm.Open(postgres.Open(ur.dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	var user User
	result := db.Where("name = ?", userName).First(&user)
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
	db, err := gorm.Open(postgres.Open(ur.dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	result := db.Model(&User{}).Where("id = ?", user.Id).Update("name", user.Name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) Delete(user u.User) error {
	db, err := gorm.Open(postgres.Open(ur.dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	result := db.Delete(&User{
		Id: user.Id.Value,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
