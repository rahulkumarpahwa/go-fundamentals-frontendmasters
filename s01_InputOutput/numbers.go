package main

import "fmt"

func numbers() {
	id := 4
	price := 45.4

	priceAsInt := int(price)
	idAsFloat := float32(id)

	fmt.Println(id, price, priceAsInt, idAsFloat)
}
