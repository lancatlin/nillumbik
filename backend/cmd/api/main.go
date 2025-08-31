package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/lancatlin/nillumbik/internal/db"
	"github.com/lancatlin/nillumbik/internal/site"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://biom:supersecretpassword@localhost:5432/nillumbik")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	querier := db.New(conn)
	r := gin.Default()
	siteCtl := site.NewController(querier)
	site.Register(r, &siteCtl)

	r.Run(":8000")
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
