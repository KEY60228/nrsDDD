package user

import (
	"database/sql"
	"fmt"
	"os"

	u "nrsDDD/domain/models/user"
)

type UserRepository struct {
	connectionString string
}

func New() (*UserRepository, error) {
	connStr := fmt.Sprintf("host=pgsql dbname=nrsDDD user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	ur := &UserRepository{
		connectionString: connStr,
	}
	return ur, nil
}

func (ur *UserRepository) Save(user u.User) error {
	db, err := sql.Open("postgres", ur.connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	return err
	// }

	_, err = db.Exec(`INSERT INTO users (id, name) VALUES ($1, $2)`, user.Id.Value, user.Name.Value)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Find(userName string) (bool, error) {
	db, err := sql.Open("postgres", ur.connectionString)
	if err != nil {
		return false, err
	}
	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	return false, err
	// }

	var cnt int32
	err = db.QueryRow(`SELECT COUNT(*) FROM users WHERE username = $1`, userName).Scan(&cnt)
	if err != nil {
		return false, err
	}

	if cnt > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
