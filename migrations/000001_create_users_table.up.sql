-- migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@pgsql:5432/$POSTGRES_DB?sslmode=disable" -path migrations up
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL
);
