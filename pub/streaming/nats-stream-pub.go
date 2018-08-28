package main

import (
	"log"
	"strconv"

	stan "github.com/nats-io/go-nats-streaming"
)

const (
	clusterID = "test-cluster"
	clientID  = "TestPub"
)

func main() {

	// connect to streaming server
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL),
	)

	if err != nil {
		log.Println(err)
		return
	}

	// ensure connection is closed at end of routine
	defer sc.Close()

	for i := 0; i < 10; i++ {
		mymessage := "Hello NATS Streaming " + strconv.Itoa(i)
		sc.Publish("foo", []byte(mymessage))
		log.Println("Published message on foo: " + mymessage)
	}
}
