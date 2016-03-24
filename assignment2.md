## Assignment 2

In this assignment, you need to persist the personal preference data that you get from the assignment 1 into a database. All the API contracts will remain the same.


### Requirement

#### Data Persistence
* All the CRUD APIs will be now storing and retrieving data from [EJDB](http://ejdb.org/doc/snippets.html#go).

#### Configuration

* Use the following [TOML](https://github.com/toml-lang/toml) (app*.toml) file to externalize all configurations. 
* Implement TOML configuration so that you can pass in TOML file from the command line.

```sh
go run app.go app1.toml
```

* app1.toml
```toml
[database]
file_name = "app1.db"

# REST API Port
port_num = 4001

[replication]
rpc_server_port_num = 3001
replica = [ "http://0.0.0.0:3002" ]
```

* app2.toml
```toml
[database]
file_name = "app2.db"

# REST API Port
port_num = 4002

[replication]
rpc_server_port_num = 3002
replica = [ "http://0.0.0.0:3001" ]
```

> [How to toml?](https://github.com/naoina/toml/tree/master/_example) 

#### Replication
 * Implement data replication for all EJDB write (Insert and Update) queries to all replica via [RPC over TCP](https://gist.github.com/jordanorelli/2629049).
 * You can do query replication either in JSON or BSON.
 * This replication must be bi-directional so that both RPC client and server's implementation is required on a server. You will be running the same code for all instances with different configuration toml files.
 * Since the main thread is listening on the REST API requests, you have to use go routine to launch RPC (server) listener. 
 * Use app1.toml for server 1 and app2.toml for another one.
 * Your code must work for any number of replica.
 
