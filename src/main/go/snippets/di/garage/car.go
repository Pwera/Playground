package garage

import (
	"strings"

	"github.com/Pwera/Playground/src/main/go/snippets/di/helpers"
)

type Car struct {
	ID    string `json:"id" bson:"_id"`
	Brand string `json:"brand" bson:"brand"`
	Color string `json:"color" bson:"color"`
}

var colorsByBrand = map[string][]string{
	"audi":    []string{"black", "white", "yellow"},
	"porsche": []string{"black", "yellow"},
	"bmw":     []string{"red", "white"},
}

func brands() []string {
	var brands []string
	for brand := range colorsByBrand {
		brands = append(brands, brand)
	}
	return brands
}

func ValidateCar(car *Car) error {
	colors, ok := colorsByBrand[car.Brand]
	if !ok {
		return helpers.NewErrValidation(
			"Brand `" + car.Brand + "` does not exist. Available brands: " +
				strings.Join(brands(), ", "))
	}

	for _, color := range colors {
		if color == car.Color {
			return nil
		}
	}

	return helpers.NewErrValidation(
		"Color `" + car.Color + "` does not exist for `" + car.Brand +
			"`. Available colors: " + strings.Join(colors, ", "),
	)
}
