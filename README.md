# GoNATSTest

Learning NATS basics with Go. Source information from which this README is derived can be found: [NATS docs](https://nats.io/documentation/)

## Prerequsites

- Your GO environment is setup and GOPATH is added to your PATH.

> brew install go

The next 2 lines you'll want to add to your ~/.bash_profile else they'll be forgotton when you restart terminal!

> export GOPATH=$HOME/workspace/go
> export PATH=$PATH:$GOPATH/bin   

Create a 'bin', 'pkg' and 'src'  directory under you GOPATH. Executables will be compiled to 'bin'. Packages will be installed in 'pkg' and all your source code and dependencies will be in 'src'

- NATS server is installed

> go get github.com/nats-io/gnatsd

navigate to gnatsd directory and run. 

> go install

This will place the gnatsd executable in your GOPATH/bin


- NATS client is installed

> go get github.com/nats-io/go-nats

- NATS streaming is installed (Not required for Basic pub-sub example)

>go get github.com/nats-io/nats-streaming-server

navigate to nats-streaming-server directory and run. 

> go install

This will place the ats-streaming-server executable in your GOPATH/bin



## Basic pub-sub

Basic 'at most once' delivery examples are found in

- sub/nats-sub.go

connects to NATS server on localhost and subscibes continuously to messages on subject 'foo'

- pub/nats-pub.go

connects to NATS server on localhost and published a single mesage on with subject 'foo' and message "Hello NATS"

### Usage

Start NATS server from terminal (Assuming you have GOPATH/bin in PATH). Note that as this is 'at most once' delivery if a subscriber isn't listening when a message is published then the message never existed as far as the subscriber is concered

>$ gnatsd

Output:

>[23563] 2018/08/24 12:12:37.753323 [INF] Starting nats-server version 1.3.0
[23563] 2018/08/24 12:12:37.753541 [INF] Git commit [not set]
[23563] 2018/08/24 12:12:37.754584 [INF] Listening for client connections on 0.0.0.0:4222
[23563] 2018/08/24 12:12:37.754610 [INF] Server is ready

Now run the subscriber in another terminal

>$ go run GOPATH/src/github.com/user/GoNATSTest/sub/nats-sub.go

Output:

>2018/08/24 12:14:26 Connected to nats://localhost:4222
2018/08/24 12:14:26 Subscribing to subject 'foo'

Now run the publisher in another terminal

>$ go run GOPATH/src/github.com/user/GoNATSTest/pub/nats-pub.go

Output:

>2018/08/24 12:15:00 Connected to nats://foo:bar@localhost:4222
2018/08/24 12:15:00 Published message on subject foo

And on the terminal for the subscriber you should now see:

>2018/08/24 12:15:00 Received message 'Hello NATS'




