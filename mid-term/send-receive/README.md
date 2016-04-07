### How to run

```sh
$ go run competingconsumers.go 
```

### Expected output

```sh
Press ENTER to exit!
Sent: ->  A
Sent: ->  B
Sent: ->  C
receiver_1  received: <-  B
receiver_2  received: <-  A
Sent: ->  D
receiver_1  received: <-  D
receiver_2  received: <-  C
Sent: ->  E
Sent: ->  F
receiver_1  received: <-  F
receiver_2  received: <-  E
Sent: ->  G
receiver_2  received: <-  G
```

> Since the receivers are acting as competing consumers, they will randomly get the data from the sender.