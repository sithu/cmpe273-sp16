package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile(fileName string)  {
    in, err := ioutil.ReadFile(fileName)
    check(err)
    
    var anyJson map[string]interface{}
    err = bson.Unmarshal(in, &anyJson)
    check(err)
    
    json, err := json.Marshal(anyJson)
    check(err)
    fmt.Println(string(json))
    
}

func grepField(file string) {
    readFile(file)
}

func main() {
    file := os.Args[1]
    grepField(file)
}
