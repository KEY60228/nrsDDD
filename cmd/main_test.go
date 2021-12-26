package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	uas "nrsDDD/application/user"
	us "nrsDDD/domain/services/user"
	testur "nrsDDD/infrastructure/testpg/user"
)

func TestMain(t *testing.T) {
	// for migration
	m, err := migrate.New(
		"file://../migrations",
		"postgres://postgresql:postgresql@testpg:5432/nrsDDD?sslmode=disable",
	)
	if err != nil {
		t.Fatal(err)
	}
	// migration up
	if err := m.Up(); err != nil {
		t.Fatal(err)
	}
	// cleanup
	t.Cleanup(func() {
		if err := m.Down(); err != nil {
			t.Fatal(err)
		}
	})

	dsn := fmt.Sprintf("host=testpg dbname=nrsDDD user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	userRepository, err := testur.New(db)
	if err != nil {
		t.Fatal(err)
	}

	userService, err := us.New(userRepository)
	if err != nil {
		t.Fatal(err)
	}

	userRegisterService, err := uas.NewUserRegisterService(userRepository, *userService, &testur.UserFactory{CurrentId: "a"})
	if err != nil {
		t.Fatal(err)
	}

	// pattern 1
	name := "KEN"

	userRegisterCommand, err := uas.NewUserRegisterCommand(name)
	if err != nil {
		t.Fatal(err)
	}

	if err := userRegisterService.Handle(*userRegisterCommand); err != nil {
		t.Fatal(err)
	}

	createdUser, err := userRepository.FindByName(name)
	if err != nil {
		t.Fatal(err)
	}
	if createdUser == nil {
		t.Fatal("ユーザーが作成されていません")
	}
	if createdUser.Name.Value != name {
		t.Fatal("ユーザー名が間違っています")
	}

	// pattern 2
	wrongName := "K"

	userRegisterCommand, err = uas.NewUserRegisterCommand(wrongName)
	if err != nil {
		t.Fatal(err)
	}

	if err := userRegisterService.Handle(*userRegisterCommand); err == nil {
		t.Fatal("ユーザー名長さチェック漏れ")
	}

	createdUser, err = userRepository.FindByName(wrongName)
	if err != nil {
		t.Fatal(err)
	}
	if createdUser != nil {
		t.Fatal("ユーザーが作成されてしまっています")
	}
}
