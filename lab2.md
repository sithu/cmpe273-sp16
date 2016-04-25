## Lab 2 

The purpose of this lab is to learn how to use consistent hashing (CH), a
client-side sharding, in a distributed datastore.


### Server side

Implement the following three REST APIs for the datastore.

* [1] Add an entry to the datastore.

***Request***

__URL: PUT /{key}/{value}__

***Response***
```sh
HTTP/1.1 204 No Content
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
``` 

---
* [2] Retrieve an entry from the datastore.

***Request***

__URL: GET /{key}__

***Response***
```sh
HTTP/1.1 200 OK
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
``` 
```json
{
   "key": 123,
   "value" : "xxxxxx"
}
``` 

---
* [3] Retrieve all entries from the datastore.

***Request***

__URL: GET /__

*** Response ***
```sh
HTTP/1.1 200 OK
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
``` 
```json
{
    [   
        {
            "key": 123,
            "value" : "xxxxxx"
        },
        {
            "key": 456,
            "value" : "xxxxxx"
        },
    ]   
}
``` 

### Client side

Implement a consistent hashing client to shard data into replica of datastore.
You can use [this simple CH implementation](https://github.com/clohfink/RendezvousHash/blob/master/src/main/java/com/csforge/ConsistentHash.java) as a reference. 


### Testing

* [1] Launch 5 instances of datastore on port 3001-3005.

```sh
go run sever.go 3001-3005
```

* [2] Run the CH client and pass the data to be sharded across the servers running on localhost's ports(3001-3005).

```sh
# {key}->{value}
go run client.go 3001-3005 1->A,2->B,3->C,4->D,5->E
```

* [3] Check the result.

```sh
curl -i "http://localhost:3001/" &&
curl -i "http://localhost:3002/" &&
curl -i "http://localhost:3003/" &&
curl -i "http://localhost:3004/" &&
curl -i "http://localhost:3005/"
```
