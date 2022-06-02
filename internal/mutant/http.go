package mutant

import (
	"github.com/gofiber/fiber/v2"
)

type (
	HttpServer struct {
		Router *fiber.App
	}

	DNAPostDto struct {
		DNA []string `json:"dna"`
	}
	DNARequestDto struct {
		IsMutant bool `json:"is_mutant"`
	}
)

func NewMutantRouter(app *fiber.App) {
	http := HttpServer{
		Router: app,
	}

	app.Post("/mutant", http.IsMutant)
}

func (h HttpServer) IsMutant(c *fiber.Ctx) error {
	dto := new(DNAPostDto)

	if err := c.BodyParser(dto); err != nil {
		return err
	}

	p := Person{DNA: DNAStructure{Nucleotides: dto.DNA}}

	if !p.IsMutant() {
		return c.Status(fiber.StatusForbidden).JSON(DNARequestDto{IsMutant: false})
	}

	return c.JSON(DNARequestDto{IsMutant: true})
}
