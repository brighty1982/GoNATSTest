package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {

	// Connect to server with auth credentials
	natsConnectionString := "nats://foo:bar@localhost:4222"
	natsConnection, _ := nats.Connect(natsConnectionString)
	defer natsConnection.Close()
	log.Println("Connected to " + natsConnectionString)

	// Publish message on subject
	subject := "foo"
	natsConnection.Publish(subject, []byte("Hello NATS"))
	log.Println("Published message on subject " + subject)
}
