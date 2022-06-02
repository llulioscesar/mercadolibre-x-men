package mutant

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Database struct {
	DB *pgx.Conn
}

func CreateDB(ctx context.Context) Database {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return Database{
		DB: conn,
	}
}

func (d *Database) Prepare(ctx context.Context) error {
	_, err := d.DB.Exec(ctx, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	_, err = d.DB.Exec(ctx, `CREATE TABLE IF NOT EXISTS dna(
		id uuid DEFAULT uuid_generate_v4(),
		nucleotides JSONB UNIQUE NOT NULL,
		is_mutant BOOLEAN NOT NULL,
		PRIMARY KEY (id)
    );`)

	_, err = d.DB.Exec(ctx, `CREATE OR REPLACE FUNCTION f_stats()
    RETURNS TABLE
            (
                j JSON
            )
    LANGUAGE plpgsql
AS
$$
BEGIN
    RETURN QUERY SELECT json_build_object('count_mutant_dna',
                                          COUNT(CASE WHEN is_mutant = TRUE THEN is_mutant END),
                                          'count_human_dna',
                                          COUNT(CASE WHEN is_mutant = FALSE THEN is_mutant END),
                                          'ratio', COUNT(CASE WHEN is_mutant = TRUE THEN is_mutant END) * 1.0 /
                                                   COUNT(CASE WHEN is_mutant = FALSE THEN is_mutant END)) j
                 FROM dna;
END;

$$;`)

	return err
}

func (d *Database) InsertDNA(ctx context.Context, dna string, isMutant bool) error {
	_, err := d.DB.Exec(ctx, `INSERT INTO dna(nucleotides, is_mutant) VALUES ($1::JSONB, $2);`, dna, isMutant)
	return err
}

func (d *Database) GetStats(ctx context.Context) (result string, err error) {
	err = d.DB.QueryRow(ctx, `SELECT * FROM f_stats();`).Scan(&result)
	return result, err
}
