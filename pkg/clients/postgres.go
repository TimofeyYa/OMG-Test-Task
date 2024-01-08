package clients

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBConfig struct {
	DataBase string
	Port     string
	User     string
	Password string
	Host     string
}

func InitPostgresClient(c DBConfig, maxAttempts int8) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.DataBase)
	return tryConnect(context.Background(), connStr, maxAttempts)
}

func tryConnect(ctx context.Context, connStr string, maxAttempts int8) (pool *pgxpool.Pool, err error) {
	for maxAttempts > 0 {
		ctxTime, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		pool, err = pgxpool.Connect(ctxTime, connStr)
		if err != nil {
			fmt.Printf("Retry connect to db:%s \n", err.Error())
			time.Sleep(5 * time.Second)
			maxAttempts--
			continue
		}
		fmt.Println("Success connect to db")
		return pool, nil
	}

	return nil, errors.New("error connect to DB")
}
