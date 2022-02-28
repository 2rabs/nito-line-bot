package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model"
	"os"
)

func NewDb(appEnv model.AppEnvironment) {
	conn, err := pgx.Connect(context.Background(), appEnv.DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
