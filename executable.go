package main

import (
	"fmt"
	"gosession/car"
)

func main() {
	carmodel := car.Car{Brand: "Nissan", Model: "Versa", Year: 2020, Price: 200000}

	carmodelv2 := car.NewCar("Kia", "Soul", 2021, 300000)

	fmt.Println(carmodel)
	fmt.Println(carmodelv2)

	list := make(map[int]car.Car)
	list[1] = car.NewCar("Chevrolet", "Spark", 2019, 50000)
	list[2] = car.NewCar("Nissan", "March", 2023, 400000)

	carlist, ok, msn := car.AddListCarMap(list, 1, car.NewCar("Kia", "Soul", 2021, 300000))
	fmt.Println(list)
	if ok {
		for key, value := range carlist {
			fmt.Printf("key [%v]  , value[%v] \n", key, value)

		}

	} else {
		fmt.Println(msn)
	}

}
