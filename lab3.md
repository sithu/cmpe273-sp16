## Lab 3 

The purpose of this lab is to learn how to use Rendezvous or Highest Random Weight (HRW) hashing, a
client-side sharding, in a distributed datastore.


### Server side

> Same as Lab 2. You can re-use the same server side implementation from the lab 2.

### Client side

Implement a HRW hashing client to shard data into replica of datastore. 
You can use [this simple Rendezvous implementation](https://github.com/clohfink/RendezvousHash/blob/master/src/main/java/com/csforge/RendezvousHash.java)  as a reference 


### Testing

* [1] Launch 5 instances of datastore on port 3001-3005.

```sh
go run sever.go 3001-3005
```

* [2] Run the HRW client and pass the data to be sharded across the servers running on localhost's ports(3001-3005).

```sh
# {key}->{value}
go run client.go 3001-3005 "1->A,2->B,3->C,4->D,5->E"
```

* [3] Check the result.

```sh
curl -i "http://localhost:3001/" &&
curl -i "http://localhost:3002/" &&
curl -i "http://localhost:3003/" &&
curl -i "http://localhost:3004/" &&
curl -i "http://localhost:3005/"
```
