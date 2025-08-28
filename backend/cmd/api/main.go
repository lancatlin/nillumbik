package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/lancatlin/nillumbik/internal/author"
	"github.com/lancatlin/nillumbik/internal/db"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:supersecretpassword@localhost:5432/nillumbik")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	querier := db.New(conn)
	r := gin.Default()
	authorCtl := author.NewController(querier)
	author.Register(r, &authorCtl)

	r.Run(":8000")
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
