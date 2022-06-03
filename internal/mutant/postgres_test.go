package mutant_test

import (
	"context"
	"encoding/json"
	"github.com/llulioscesar/mercadolibre-x-men/internal/mutant"
	"github.com/pashagolub/pgxmock"
	"testing"
)

func GetConnTest(t *testing.T) pgxmock.PgxConnIface {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectPrepare("insert_dna", `INSERT INTO dna`).
		ExpectExec().
		WithArgs(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`, true).
		WillReturnResult(pgxmock.NewResult("insert_dna", 1))

	rows := pgxmock.NewRows([]string{"f"}).
		AddRow(`{"count_mutant_dna":40,"count_human_dna":100,"ratio":0.4}`)
	mock.ExpectQuery(`SELECT (.+) FROM f_stats()`).
		WillReturnRows(rows)

	mock.ExpectExec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS dna`)
	mock.ExpectExec(`CREATE OR REPLACE FUNCTION f_stats`)

	return mock
}

func TestDatabase_InsertDNA(t *testing.T) {
	ctx := context.Background()

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mock.Close(ctx)

	dto := mutant.DNAPostDto{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}
	p := mutant.Person{DNA: mutant.DNAStructure{Nucleotides: dto.DNA}}

	raw, err := json.Marshal(&dto)
	if err != nil {
		t.Error(err)
	}

	//rows := pgxmock.NewRows([]string{"nucleotides", "is_mutant"}).
	//	AddRow(string(raw), p.IsMutant())

	mock.ExpectPrepare("insert_dna", `INSERT INTO dna`).
		ExpectExec().
		WithArgs(string(raw), p.IsMutant()).
		WillReturnResult(pgxmock.NewResult("insert_dna", 1))

	db := mutant.Database{
		DB: mock,
	}

	err = db.InsertDNA(ctx, string(raw), p.IsMutant())
	if err != nil {
		t.Errorf("error on test InsertDNA %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDatabase_GetStats(t *testing.T) {
	ctx := context.Background()

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mock.Close(ctx)

	rows := pgxmock.NewRows([]string{"f"}).
		AddRow(`{"count_mutant_dna":40,"count_human_dna":100,"ratio":0.4}`)
	mock.ExpectQuery(`SELECT (.+) FROM f_stats()`).
		WillReturnRows(rows)

	db := mutant.Database{
		DB: mock,
	}

	raw, err := db.GetStats(ctx)
	if err != nil {
		t.Errorf("error on test GetStats %s", err)
	}

	t.Log(raw)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDatabase_Prepare(t *testing.T) {
	ctx := context.Background()

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mock.Close(ctx)

	mock.ExpectExec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS dna`)
	mock.ExpectExec(`CREATE OR REPLACE FUNCTION f_stats`)

	db := mutant.Database{
		DB: mock,
	}

	db.Prepare(ctx)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
