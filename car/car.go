package car

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
	Price float32
}

func NewCar(brand, model string, year int, price float32) Car {
	return Car{Brand: brand, Model: model, Year: year, Price: price}
}

func AddListCarMap(list map[int]Car, idKey int, item Car) (map[int]Car, bool, string) {
	if _, ok := list[idKey]; ok {
		return list, false, fmt.Sprintf("Car with id %v already exists", idKey)
	}

	list[idKey] = item
	return list, true, "Success"
}
