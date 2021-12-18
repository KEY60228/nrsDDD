package application

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"nrsDDD/Domain/Models/user"
	userService "nrsDDD/Domain/Services/user"
)

type Program struct{}

func (p *Program) CreateUser(userName string) (*user.User, error) {
	exists, err := userService.Exists(userName)
	if exists {
		return nil, fmt.Errorf("%vは既に存在しています", userName)
	}

	u, err := user.New(userName)
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("host=pgsql dbname=nrsDDD user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }

	_, err = db.Exec(`INSERT INTO users (id, name) VALUES ($1, $2)`, u.Id.Value, u.Name.Value)
	if err != nil {
		return nil, err
	}

	return u, nil
}
