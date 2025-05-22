package config

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"sync"
	"github.com/joho/godotenv"
	"os"
)
var Conn *pgxpool.Pool
var startApp sync.Once

func Connectdb(){
	startApp.Do(func() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	url := os.Getenv("DB_URL")
	cfg, _ := pgxpool.ParseConfig(url)
	//cfg.ConnConfig.StatementCacheCapacity = 0
	conn, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		fmt.Println(err)
	}
	Conn = conn
	fmt.Println("PING", Conn.Ping(context.Background()))
	})
}