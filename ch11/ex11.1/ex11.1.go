package main

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House

	house.Address = "123 Main St"
	house.Size = 2000
	house.Price = 350000.00
	house.Type = "Single Family"

	println("Address:", house.Address)
	println("Size:", house.Size)
	println("Price:", house.Price)
	println("Type:", house.Type)
}
