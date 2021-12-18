package application

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	u "nrsDDD/domain/models/user"
	us "nrsDDD/domain/services/user"
)

type Program struct{}

func (p *Program) CreateUser(userName string) (*u.User, error) {
	exists, err := us.Exists(userName)
	if exists {
		return nil, fmt.Errorf("%vは既に存在しています", userName)
	}

	user, err := u.New(userName)
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

	_, err = db.Exec(`INSERT INTO users (id, name) VALUES ($1, $2)`, user.Id.Value, user.Name.Value)
	if err != nil {
		return nil, err
	}

	return user, nil
}
