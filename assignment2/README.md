### Install dependencies
```sh
go get github.com/drone/routes

# For other OSes, see @http://wiki.bazaar.canonical.com/Download
brew install bzr

# BSON Lib
go get labix.org/v2/mgo/bson

# EJDB
go get github.com/mkilling/goejdb
```

### Run the app

```sh
go run app.go
````

Now, open this URL (http://localhost:3000/profile/foo@gmail.com) in a web browser!


[Routes]:https://github.com/drone/routes