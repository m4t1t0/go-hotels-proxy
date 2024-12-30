package home

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"os"
)

func Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		connStr := fmt.Sprintf(
			"postgres://%s:%s@%s:5432/%s",
			os.Getenv("POSTGRESQL_READS_USERNAME"),
			os.Getenv("POSTGRESQL_READS_PASSWORD"),
			os.Getenv("POSTGRESQL_READS_HOST"),
			os.Getenv("POSTGRESQL_READS_DBNAME"),
		)

		conn, err := pgx.Connect(context.Background(), connStr)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}

		defer conn.Close(context.Background())

		type Accommodation struct {
			Id       string
			Name     string
			Category int
		}

		rows, _ := conn.Query(
			context.Background(),
			"SELECT id, name, category FROM main_reads.all_accommodations aa LIMIT 10",
		)

		var accommodations []Accommodation

		for rows.Next() {
			var acc Accommodation
			err := rows.Scan(&acc.Id, &acc.Name, &acc.Category)
			if err != nil {
				return c.SendString(fmt.Sprintf("Error Fetching Accommodation Details: %v\n", err))
			}
			accommodations = append(accommodations, acc)
		}

		return c.SendString(fmt.Sprintf("Accommodations\n\n%v", accommodations))
	}
}
