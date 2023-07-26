package connect

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DbConection() {
	var err error

	DbUrl := "postgres://postgres:MF13012003@localhost:5432/db_personal_web"

	Conn, err = pgx.Connect(context.Background(), DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

}
