package main

import (
    "fmt"
    "os"
    "time"
    "io/ioutil"
    "gopkg.in/mgo.v2/bson"
    "github.com/jasonlvhit/gocron"
)

type Person struct {
    Name string
    Phone string ",omitempty"
}
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func writeFile(data []byte, fileName string)  {
    file, err := os.Create(fileName)
    check(err)
    
    defer file.Close()   
    
    numBytes, err := file.Write(data)
    fmt.Printf("\nwrote %d bytes to %s\n", numBytes, fileName)
    
    file.Sync()
}

func readFile(fileName string)  {
    in, err := ioutil.ReadFile(fileName)
    check(err)
    
    person := new(Person)
    bson.Unmarshal(in, &person)
    check(err)
    
    fmt.Printf("Read from BSON file: Person.Name=%s\n", person.Name)
}

func createBSON() []byte {
    data, err := bson.Marshal(&Person{Name:"Bob"})
    check(err)
    fmt.Printf("%q", data)
    
    return data    
}

func makeTimestamp() int64 {
    return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}

func task() {
    t := makeTimestamp()
    fmt.Printf("Running task...@%d\n", t)
    data := createBSON()
    writeFile(data, "data.bson")
    readFile("data.bson")
}

func main() {
    s := gocron.NewScheduler()
    s.Every(5).Seconds().Do(task)
    <- s.Start()
}
