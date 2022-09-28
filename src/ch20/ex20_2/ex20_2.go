package main

import (
	"golangstudy/ch20/ex20_2/sender/fedex"
	"golangstudy/ch20/ex20_2/sender/koreapost"
)

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

type Sender interface {
	Send(parcel string)
}

func main() {
	sender1 := &fedex.FedexSender{}
	SendBook("어린 왕자", sender1)
	SendBook("그리스인 조르바", sender1)
	sender2 := &koreapost.PostSender{}
	SendBook("어린 왕자", sender2)
	SendBook("그리스인 조르바", sender2)
}
