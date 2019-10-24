package garage

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
	// color, ok := colorsByBrand[car.Brand]
	// if !ok{
	// 	return helpers.NewErrValidation()
	// }

	return nil
}
