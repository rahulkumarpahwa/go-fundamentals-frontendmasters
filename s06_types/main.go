package main

import "fmt"

type integer = int            // alias
type json = map[string]string //alias
var x integer                 // usage of the type with alias

// type :
type distance float64 // now here distance is the new type with value of the float64. this will also create the convertor of the same type as well which will conevrt to the float64 (for here one example)
// also this one type creates the semantics(logic) for the values we put with this type. like it can be used to show the distance in miles as it has the type float64.

type distanceKm float64 // represents the km

func test() {
	d := distance(6.4)
	fmt.Println(d)
}

// another example of the semantics with type as:
func ToKm(miles distance) distanceKm { // this one converts the miles to km.
	return distanceKm(1.60934 * miles)
}

// making a method connected to a type. here we are not creating the method to the class but the method to a type. this method Tokm2 will take the distance in distance type and return to km in distanceLKm type.
// this is similar in js adding method to the prototype.
func (miles distance) Tokm2() distanceKm {
	return distanceKm(1.60934 * miles)
}

// similar method to miles in type distanceKm
func (km distanceKm) ToMiles() distance {
	return distance(0.62137119 * km)
}

func main() {
	d := distance(6.4)
	fmt.Println(d)
	distanceInKm := d.Tokm2() // this converts the miles to km.
	fmt.Println(distanceInKm)

	fmt.Println("Distance in Miles from km : ", distanceKm.ToMiles(distanceInKm))
	// or :
	fmt.Println("Distance in Miles from km : ", distanceInKm.ToMiles())

	// so, here we learn that we can create the method for the type. this will add semantic meaning and create more descriptive types with context
}
