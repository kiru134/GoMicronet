package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name of the excahnge
		"topic",      //type
		false,        //durable?
		false,        //auto-deleted?
		false,        //internal
		false,        //no wait
		nil,          //arguments?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    //name- if the queue name is empty the server vl return a unique queue name
		false, //durable?
		false, // delete when unsused?
		true,  //exclusive?
		false, //no wait?
		nil,   //arguments?
	)
}
