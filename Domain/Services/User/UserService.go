package user

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Exists(userName string) (bool, error) {
	// 重複を確認する
	connStr := fmt.Sprintf("host=pgsql dbname=nrsDDD user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return false, err
	}
	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
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
