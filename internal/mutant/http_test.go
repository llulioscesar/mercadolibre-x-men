package mutant_test

import (
	"bytes"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/llulioscesar/mercadolibre-x-men/internal/mutant"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type HTTPTest struct {
	description  string
	route        string
	expectedCode int
}

func TestHttpServer_IsMutant(t *testing.T) {
	test := HTTPTest{
		description:  "Validate mutant post HTTP status 200",
		route:        "/mutant",
		expectedCode: 200,
	}

	ctx := context.Background()
	app := fiber.New()

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := mutant.Database{
		DB: mock,
	}

	defer mock.Close(ctx)

	mock.ExpectPrepare("insert_dna", `INSERT INTO dna`).
		ExpectExec().
		WithArgs(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`, true).
		WillReturnResult(pgxmock.NewResult("insert_dna", 1))

	server := mutant.HttpServer{
		Router:  app,
		DB:      db,
		Context: ctx,
	}

	app.Post(test.route, server.IsMutant)

	req := httptest.NewRequest("POST", test.route, bytes.NewReader([]byte(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, 1)

	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
}

func TestHttpServer_IsMutant2(t *testing.T) {
	test := HTTPTest{
		description:  "Validate mutant post HTTP status 200",
		route:        "/mutant",
		expectedCode: 403,
	}

	ctx := context.Background()
	app := fiber.New()

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := mutant.Database{
		DB: mock,
	}

	defer mock.Close(ctx)

	mock.ExpectPrepare("insert_dna", `INSERT INTO dna`).
		ExpectExec().
		WithArgs(`{"dna":["ATGCGA","CTGTGC","TTATGT","AGAAGG","CGCCTA","TCACTG"]}`, false).
		WillReturnResult(pgxmock.NewResult("insert_dna", 1))

	server := mutant.HttpServer{
		Router:  app,
		DB:      db,
		Context: ctx,
	}

	app.Post(test.route, server.IsMutant)

	req := httptest.NewRequest("POST", test.route, bytes.NewReader([]byte(`{"dna":["ATGCGA","CTGTGC","TTATGT","AGAAGG","CGCCTA","TCACTG"]}`)))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, 1)

	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
}

func TestHttpServer_Stats(t *testing.T) {
	test := HTTPTest{
		description:  "Get Stats HTTP status 200",
		route:        "/stats",
		expectedCode: 200,
	}

	ctx := context.Background()
	app := fiber.New()

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

	server := mutant.HttpServer{
		Router:  app,
		DB:      db,
		Context: ctx,
	}

	app.Get(test.route, server.Stats)

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := app.Test(req, 1)
	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
}
