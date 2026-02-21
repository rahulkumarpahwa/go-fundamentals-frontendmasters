package datatypes

// here we are making the struct of the data of currency we get and it is called here the Rate. this will be exported so it is capital letter.
type Rate struct {
	Currency string
	Price float64 // here we can use the float32 as well if we need to two decimal in the value but mind that the json encoder and decoder in go provides the data in float64 format, which is more precise and efficient. Also, we need to convert the type to float64 everytime we send or get data.
}