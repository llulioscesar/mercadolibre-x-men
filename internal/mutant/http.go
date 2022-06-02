package mutant

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type (
	HttpServer struct {
		Router  *fiber.App
		DB      Database
		Context context.Context
	}

	DNAPostDto struct {
		DNA []string `json:"dna"`
	}
	DNAResponseDto struct {
		IsMutant bool `json:"is_mutant"`
	}
	StatsResponseDto struct {
		Human  float64 `json:"count_human_dna"`
		Mutant float64 `json:"count_mutant_dna"`
		Ratio  float64 `json:"ratio"`
	}
)

func NewMutantRouter(ctx context.Context, app *fiber.App, db Database) {

	err := db.Prepare(ctx)
	if err != nil {
		panic(err)
	}

	http := HttpServer{
		Router:  app,
		DB:      db,
		Context: ctx,
	}

	app.Post("/mutant", http.IsMutant)
	app.Get("/stats", http.Stats)
}

func (h HttpServer) Stats(c *fiber.Ctx) error {
	var result StatsResponseDto
	raw, err := h.DB.GetStats(h.Context)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(raw), &result)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h HttpServer) IsMutant(c *fiber.Ctx) error {
	dto := new(DNAPostDto)

	if err := c.BodyParser(dto); err != nil {
		return err
	}

	p := Person{DNA: DNAStructure{Nucleotides: dto.DNA}}
	isMutant := p.IsMutant()

	b, err := json.Marshal(&dto)
	if err != nil {
		return err
	}

	err = h.DB.InsertDNA(h.Context, string(b), isMutant)
	if err != nil {

	}

	if !isMutant {
		return c.Status(fiber.StatusForbidden).JSON(DNAResponseDto{IsMutant: false})
	}

	return c.JSON(DNAResponseDto{IsMutant: true})
}
