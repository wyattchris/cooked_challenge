package handlers;

import (
	"github.com/gofiber/fiber/v2"
	"github.com/GenerateNU/cooked/backend/internal/errs"
	"github.com/GenerateNU/cooked/backend/internal/types"

	"github.com/google/uuid"
	"fmt"
)


type CreateRecipeReqeust struct {
	Name         string    `json:"name"`
	Cook         types.Duration  `json:"cook_duration"` 
	Instructions string    `json:"instructions"`
	ImageURL     types.URL       `json:"image_url"`
	Meal         types.Meal      `json:"meal"`
}

func (s *Service) CreateRecipe(c *fiber.Ctx) error {

	var req CreateRecipeReqeust

	if err := c.BodyParser(&req); err != nil {
		fmt.Println(err)
		return errs.InvalidJSON()
	}

	recipe, err := s.store.CreateRecipe(c.Context(), req.into())
	if err != nil {
		return errs.InternalServerError()
	}

	return c.Status(fiber.StatusCreated).JSON(recipe)
}

func (r *CreateRecipeReqeust) into() types.Recipe {
	return types.Recipe{
		ID: 		  uuid.New(),
		Name:         r.Name,
		Cook: 			r.Cook,
		Instructions: r.Instructions,
		ImageURL:     r.ImageURL,
		Meal:         r.Meal,
	}
}
