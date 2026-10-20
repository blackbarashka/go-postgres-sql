package data

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	//change.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:admin@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	err_ := conn.Ping(ctx)
	if err_ != nil {
		log.Fatal(err_)
	}
	fmt.Println("Successfully connected to postgres")
}
