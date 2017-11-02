package migrations

import (
	"database/sql"

	"fmt"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
)

func Up(dbname string, dialect string, migrationsDir string) {
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", dbname)
	db, err := sql.Open(dialect, dbinfo)
	if err != nil {
		fmt.Printf("sql.Open failed due to: %s", err.Error())
	}
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	migrate.Exec(db, dialect, migrations, migrate.Up)
}

func Down(dbname string, dialect string, migrationsDir string) {
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", dbname)
	db, err := sql.Open(dialect, dbinfo)
	if err != nil {
		fmt.Printf("sql.Open failed due to: %s", err.Error())
	}
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	migrate.Exec(db, dialect, migrations, migrate.Down)
}
