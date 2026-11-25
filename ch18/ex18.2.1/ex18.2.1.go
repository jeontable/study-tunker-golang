package main

import (
	"github.com/tuckersGo/musthaveGo2/ch18/fedex"
	"github.com/tuckersGo/musthaveGo2/ch18/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.PostSender{}
	fedexSender := &fedex.FedexSender{}

	SendBook("어린 왕자", koreaPostSender)
	SendBook("어린 왕자", fedexSender)
}
