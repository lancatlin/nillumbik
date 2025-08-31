package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/lancatlin/nillumbik/internal/db"
	"github.com/lancatlin/nillumbik/internal/site"
	"github.com/lancatlin/nillumbik/internal/species"
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

	speciesCtl := species.NewController(querier)
	species.Register(r, &speciesCtl)

	r.Run(":8000")
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
