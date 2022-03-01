package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model"
	"os"
)

func NewDb(appEnv *model.AppEnvironment) {
	conn, err := pgxpool.Connect(context.Background(), appEnv.DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	var memberId int64
	var nickname string
	err = conn.QueryRow(
		context.Background(),
		"select member_id, nickname from member where member_id=$1",
		1,
	).Scan(&memberId, &nickname)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(memberId, nickname)
}
