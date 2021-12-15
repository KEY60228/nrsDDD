-- migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@pgsql:5432/$POSTGRES_DB?sslmode=disable" -path migrations down
DROP TABLE IF EXISTS users;
