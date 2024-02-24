package main

import (
	"cocom/internal/api"
	"cocom/internal/machine"
	"cocom/internal/models"
	"log/slog"
	"net/http"
)

func main() {

	sodas := []*models.Soda{
		&models.Soda{
			ID:          "fizz",
			Name:        "Fizz",
			Description: "An effervescent fruity experience with hints of grape and coriander.",
			Price:       100,
			Quantity:    100,
		},
		&models.Soda{
			ID:          "pop",
			Name:        "Pop",
			Description: "An explosion of flavor that will knock your socks off!",
			Price:       100,
			Quantity:    100,
		},
		&models.Soda{
			ID:          "cola",
			Name:        "Cola",
			Description: "A basic no nonsense cola that is the perfect pick me up for any occasion.",
			Price:       100,
			Quantity:    200,
		},
		&models.Soda{
			ID:          "mega-pop",
			Name:        "Mega Pop",
			Description: "Not for the faint of heart.  So flavorful and so invigorating, it should probably be illegal.",
			Price:       100,
			Quantity:    50,
		},
	}

	vm := machine.NewMachine(sodas)

	handler := api.NewHandler(vm)

	handler.RegisterRoutes()

	slog.Debug("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error(err.Error())
	}
}
