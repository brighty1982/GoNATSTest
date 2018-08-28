# Go and NATS tutorial

Learning NATS basics with Go. Source information from which this README is derived can be found: [NATS docs](https://nats.io/documentation/)

## Prerequsites

- Your GO environment is setup and ```GOPATH``` is added to your ```PATH```. You can download the binary from [Go install](https://golang.org/dl/) or use Homebrew as below.

```
$ brew install go
```

The next 2 lines you'll want to add to your ```~/.bash_profile``` else they'll be forgotton when you restart terminal!

```
$ export GOPATH=$HOME/workspace/go

$ export PATH=$PATH:$GOPATH/bin
```   

Create a ```'bin', 'pkg' and 'src'```  directory under you ```GOPATH```. Executables will be compiled to ```'bin'```. Packages will be installed in ```'pkg'``` and all your source code and dependencies will be in ```'src'```

- NATS server is installed
```
$ go get github.com/nats-io/gnatsd
```
navigate to gnatsd directory and run. 
```
$ go install
```
This will place the gnatsd executable in your ```GOPATH/bin```


- NATS client is installed
```
$ go get github.com/nats-io/go-nats
```
- NATS streaming server is installed (Not required for Basic pub-sub example)
```
$ go get github.com/nats-io/nats-streaming-server
```
navigate to ```nats-streaming-server directory``` and run. 
```
$ go install
```
This will place the nats-streaming-server executable in your ```GOPATH/bin```

- NATS streaming client is installed (Not required for Basic pub-sub example)
```
$ go get github.com/nats-io/go-nats-streaming
```


## Basic pub-sub

Basic 'at most once' delivery examples are found in

- ```sub/nats-sub.go```

connects to NATS server on localhost and subscibes continuously to messages on subject 'foo'

- ```pub/nats-pub.go```

connects to NATS server on localhost and publishes a single message on with subject ```'foo'``` and message ```"Hello NATS"```

### Usage

Start NATS server from terminal (Assuming you have ```GOPATH/bin``` in ```PATH```). Note that as this is 'at most once' delivery if a subscriber isn't listening when a message is published then the message never existed as far as the subscriber is concered
```
$ gnatsd
```
Output:
```
[23563] 2018/08/24 12:12:37.753323 [INF] Starting nats-server version 1.3.0
[23563] 2018/08/24 12:12:37.753541 [INF] Git commit [not set]
[23563] 2018/08/24 12:12:37.754584 [INF] Listening for client connections on 0.0.0.0:4222
[23563] 2018/08/24 12:12:37.754610 [INF] Server is ready
```

Now run the subscriber in another terminal
```
$ go run GOPATH/src/github.com/user/GoNATSTest/sub/nats-sub.go
```

Output:
```
2018/08/24 12:14:26 Connected to nats://localhost:4222
2018/08/24 12:14:26 Subscribing to subject 'foo'
```
Now run the publisher in another terminal
```
$ go run GOPATH/src/github.com/user/GoNATSTest/pub/nats-pub.go
```
Output:
```
2018/08/24 12:15:00 Connected to nats://foo:bar@localhost:4222
2018/08/24 12:15:00 Published message on subject foo
```
And on the terminal for the subscriber you should now see:
```
2018/08/24 12:15:00 Received message 'Hello NATS'
```

## Basic Streaming pub-sub

Basic 'at least once' delivery examples are found in

- ```sub/streaming/nats-stream-sub.go```

connects to NATS streaming server on localhost and subscibes continuously to messages on subject 'foo' with the 'durable' flag set on the connection to ensure it will always start consuming from the last message that it acknowledged to the server

- ```pub/streaming/nats-stream-pub.go```

connects to NATS streaming server on localhost and publishes 10 messages on with subject 'foo' and message "Hello NATS " + i.

### Usage

Start NATS streaming server from terminal (Assuming you have ```GOPATH/bin``` in ```PATH```). Note that as this is 'at least once' delivery a message published will still be picked up by the subscriber even if the client isn't running when it is published. What won't be guarenteed is that a client will only receive a message 'once'. This is the preferred method of using NATS as it ensure messages won't be lost when a client subsciber service goes down.
```
$ nats-streaming-server
```
Output:
```
$ [31661] 2018/08/28 09:20:44.547304 [INF] STREAM: Starting nats-streaming-server[test-cluster] version 0.11.0
$ [31661] 2018/08/28 09:20:44.547438 [INF] STREAM: ServerID: 2X1raoN93a8mWfH0jGnNGd
$ [31661] 2018/08/28 09:20:44.547442 [INF] STREAM: Go version: go1.10.3
$ [31661] 2018/08/28 09:20:44.548048 [INF] Starting nats-server version 1.3.0
$ [31661] 2018/08/28 09:20:44.548129 [INF] Git commit [not set]
$ [31661] 2018/08/28 09:20:44.548388 [INF] Listening for client connections on 0.0.0.0:4222
$ [31661] 2018/08/28 09:20:44.548397 [INF] Server is ready
$ [31661] 2018/08/28 09:20:44.576881 [INF] STREAM: Recovering the state...
$ [31661] 2018/08/28 09:20:44.576906 [INF] STREAM: No recovered state
$ [31661] 2018/08/28 09:20:44.832685 [INF] STREAM: Message store is MEMORY
$ [31661] 2018/08/28 09:20:44.832836 [INF] STREAM: ---------- Store Limits ----------
$ [31661] 2018/08/28 09:20:44.832856 [INF] STREAM: Channels:                  100 *
$ [31661] 2018/08/28 09:20:44.832863 [INF] STREAM: --------- Channels Limits --------
$ [31661] 2018/08/28 09:20:44.832872 [INF] STREAM:   Subscriptions:          1000 *
```

Now run the publisher in another terminal. Note that in the previous example we ran the subscriber first to ensure it was listening when the messages were sent. This time we don't have to as the client will get them anyway.

```
$ go run GOPATH/src/github.com/user/GoNATSTest/pub/streaming/nats-stream-pub.go
```
Output:

```
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 0
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 1
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 2
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 3
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 4
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 5
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 6
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 7
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 8
$ 2018/08/28 09:47:07 Published message on foo: Hello NATS Streaming 9
```


Now run the subscriber in another terminal
```
$ go run GOPATH/src/github.com/user/GoNATSTest/sub/streaming/nats-stream-sub.go
```
Output:

```
$ 2018/08/28 10:29:44 Subscribing to subject 'foo'
$ Received a message:  Hello NATS Streaming 0
$ Received a message:  Hello NATS Streaming 1
$ Received a message:  Hello NATS Streaming 2
$ Received a message:  Hello NATS Streaming 3
$ Received a message:  Hello NATS Streaming 4
$ Received a message:  Hello NATS Streaming 5
$ Received a message:  Hello NATS Streaming 6
$ Received a message:  Hello NATS Streaming 7
$ Received a message:  Hello NATS Streaming 8
$ Received a message:  Hello NATS Streaming 9
```

Even though the subsciber wasn't running when the messages were sent it still received them.

If you end the subscriber process ``` control 'c' ``` and then run it again it will not process the messages again. This is because the ```durable``` flag specified in the connection. Doing this causes the NATS Streaming server to track the last acknowledged message for that clientID + durable name, so that only messages since the last acknowledged message will be delivered to the client.

You can play with the various subscriber replay options here: [NATS client replay](https://github.com/nats-io/go-nats-streaming)