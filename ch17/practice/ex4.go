package main

import "time"

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt           *Product
	ShippingTime  time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	parcel := Parcel{Pdt: p, ShippingTime: time.Now()}
	return &parcel
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}
