## Assignment 2

In this assignment, you need to persist the personal preference data that you get from the assignment 1 into a database. All the API contracts will remain the same.


### Requirement

#### Data Persistence
* You will be storing the personal reference data into [EJDB](http://ejdb.org/doc/snippets.html#go). All the CRUD APIs will now save and fetch data from EJDB.

#### Configuration

* Use the following [TOML](https://github.com/toml-lang/toml) (app.toml) file to externalize all configurations. 

* app.toml
```toml
[database]
file_name = "app.db"

port_num = 3000
```

> [How to toml?](https://github.com/naoina/toml/tree/master/_example) 

