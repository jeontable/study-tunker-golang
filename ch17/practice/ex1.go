package main

type Cart struct {
	productList string
}

func (c *Cart) AddProduct(product string) {
	if c.productList != "" {
		c.productList += ", "
	}
	c.productList += product
}
