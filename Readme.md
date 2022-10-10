# An implementation "Word of Wisdom" tcp server with DDoS protection based on a proof of work algorithm

[![Actions Status](https://github.com/nightlord189/tcp-pow-go/workflows/main/badge.svg)](https://github.com/nightlord189/tcp-pow-go/actions)

## Requirements

- TCP-server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work),
  the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Prof Of Work verification, server should send one of the quotes from "word of wisdom" book or any other
  collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge.

## Pre-setup

- [Go 1.19+](https://go.dev/dl/) installed (to run the tests, the server or the client and check the code by linters)

## Run the server

```shell
make run-server
```

It runs the server with `race` golang flag and default configuration file `config.json`.

## Run the client

```shell
make run-server
```

It runs the client with `race` golang flag and default configuration file `config.json`.

## Launch unit-tests:

```
make test
```

It launches all available unit-tests over the repository.

## Launch linters code check:

```
make linter
```

It launches golangci-lint code validation.

## 4. Protocol description

The solution uses TCP-based as protocol.
For messaging I used GOB (golang native binary format) with the following design:

```go
type Message struct {
  Header Header
  Body   any
}
```

Where `Header` is one go:

- AskType - the client asks new challenge from the server
- ChallengeType - the server message with challenge structure that is needed to compute
- AnswerType - the client message with the solved challenge
- GrandType - the server response with a quote of Word of Wisdom
- ErrorType - the server response for unsuccessful cases

`Body` is an optional value with payload structure (it depends on the header type).

[There](./protocol.puml) is located a sequence diagram of the protocol

## 5. Proof of Work algorithm

Idea of Proof of Work to protect DDoS is based on the follow solution:

- before sending some data to a client has to solve mathematical challenge that requires sensitive value of computer
  power
- after it sends payload and the result of the challenge
- a server receives the request and validates the result of the challenge
- if the result is ok then it handles the payload
- main idea of the challenge is to require more computational work on client-side and challenge's result verification
  requires much
  less on the server-side

### 5.1 PoW algorithm choose

There are different algorithms Proof of Work:

- [Hashcash](https://en.wikipedia.org/wiki/Hashcash)
- [Guided tour puzzle](https://en.wikipedia.org/wiki/Guided_tour_puzzle_protocol)
- [Merkle tree](https://en.wikipedia.org/wiki/Merkle_tree)

I compared these three algorithms as more famous and having most extensive description.

After comparison, I made a chose in favor of Hashcash.

There are lists of the algorithms pros and cons:

**Hashcash pros**

- easy to implement
- simplicity of validation on server side
- tons of documentation
- ability to dynamically manage complexity for client
- I used it before several times

**Hashcash cons**

- time computing is depends on power of client's machine. Very weak client's machine may not solve challenge for TTL. At
  the same time more powerful machines can implement some kind of DDoS.

**Merkle tree pros**

- the solution is stronger comparing with Hashcash

**Merkle tree cons**

- a server should makes too much work to validate a client's solution. For tree consists of 4 leaves and 3 depth a
  server
  will spend 3 hash calculations.

**Guided tour puzzle cons**

- speaking of the requirements the PoW isn't good solution because a client should regularly request a server about the
  next parts of guide that complicates logic of the protocol

But all of these cons could be solved in a real application.

## Structure of the project

- cmd/client - main.go for the client
- cmd/server - main.go for the server
- internal/config - the parser of a configuration file
- internal/hashcacsh - PoW algorithm
- internal/protocol - contains structures, encoder and decoder for TCP-protocol messaging
- internal/quotes - contains Word of Wisdom quotes
- pkg/client - the client logic
- pkg/server - the server logic

## Ways to improve

Of course, every project could be improved. This project also has some ways to improve:

- add context and time-to-live for every connection
- cover integration test
- add multiclient support (current solution is sync)
- add on-fly the server/client configuration (to improve DDoS protection flexibly)
- migrate in-memory quotes to SQL database like PostgreSQL
- add session tokens to handle series of requests
- add Docker configurations

## Outs examples

**The server log:**

```shell
2022/10/10 16:12:39 new client: 127.0.0.1:57014
2022/10/10 16:12:42 new client: 127.0.0.1:57015
2022/10/10 16:12:45 new client: 127.0.0.1:57016
2022/10/10 16:12:48 new client: 127.0.0.1:57042
2022/10/10 16:12:51 new client: 127.0.0.1:57043
```

**The client log:**
```shell
2022/10/10 16:12:39 connected to 127.0.0.1:8080
quote: Good people are good because they've come to wisdom through failure. We get very little wisdom from success, you know.
2022/10/10 16:12:42 connected to 127.0.0.1:8080
quote: Each morning when I open my eyes I say to myself: I, not events, have the power to make me happy or unhappy today. I can choose which it shall be. Yesterday is dead, tomorrow hasn't arrived yet. I have just one day, today, and I'm going to be happy in it.
2022/10/10 16:12:45 connected to 127.0.0.1:8080
quote: If any of you lacks wisdom, let him ask of God, who gives to all liberally and without reproach, and it will be given to him.
2022/10/10 16:12:48 connected to 127.0.0.1:8080
quote: Marriage is about compromise. Being a winner isn’t important, working together is.
2022/10/10 16:12:51 connected to 127.0.0.1:8080
quote: Marriage takes work. But the payoff is years of happiness.
```

## Feedback

Before developing of the solution I made decision to implement it within minimum test cycles of  start-stop.

The first success run of the server and the client were after just 7th attempts.

It was my personal challenge for fun :)