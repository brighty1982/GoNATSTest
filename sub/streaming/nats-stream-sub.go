package main

import (
	"fmt"
	"log"
	"runtime"

	stan "github.com/nats-io/go-nats-streaming"
)

const (
	clusterID = "test-cluster"
	clientID  = "TestSub"
	durableID = "myDurableID"
)

func main() {

	// connect to streaming server on local host
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL),
	)

	if err != nil {
		log.Println(err)
		return
	}

	// Subscribe to subject
	log.Println("Subscribing to subject 'foo'")

	// Replay all messages for subject
	// that haven't been 'acked' before
	sc.Subscribe("foo", func(m *stan.Msg) {
		m.Ack() // Manual ACK
		fmt.Println("Received a message: ", string(m.Data))
	}, stan.DurableName(durableID))
	//, stan.DurableName(durableID), stan.DeliverAllAvailable())
	runtime.Goexit()

}
