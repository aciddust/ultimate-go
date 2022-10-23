package main

import (
	"gobyexample/awesome-go/rabbitmq"
)

func main() {
	exchange := "amp.fanout"
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	err = ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil)
	if err != nil {
		return
	}
	go mq.Consumer(ch, exchange, "c1", "", "")
	go mq.Consumer(ch, exchange, "c2", "", "")
	for {

	}
}
