package main

import (
    "fmt"
    "os"
    "time"
    "io/ioutil"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type Request struct {
    Url string `bson:"url,omitempty" json:"url,omitempty"`
    Method string `bson:"method,omitempty" json:"method,omitempty"`
    HttpHeaders map[string]string `bson:"http_headers,omitempty" json:"http_headers,omitempty"`
    Body map[string]string `bson:"body,omitempty" json:"body,omitempty"`
}

type Input struct {
    Id int64 `bson:"id,omitempty" json:"id"`
    SuccessHttpResponseCode int `bson:"success_http_response_code,omitempty" json:"success_http_response_code"`
    MaxRetries int `bson:"max_retries,omitempty" json:"max_retries,omitempty"`
    CallbackWebhookUrl string `bson:"callback_webhook_url,omitempty" json:"callback_webhook_url,omitempty"`
    Request Request `bson:"request,omitempty" json:"request,omitempty"`
}

type InputArray struct {
    Inputs []Input `bson:"inputs" json:"inputs"`
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
    fmt.Printf("\nWrote %d bytes to %s\n", numBytes, fileName)
    
    err = file.Sync()
    check(err)
}

func readFile(fileName string)  {
    in, err := ioutil.ReadFile(fileName)
    check(err)
    
    var anyJson map[string]interface{}
    err = bson.Unmarshal(in, &anyJson)
    
    json, err := json.Marshal(anyJson)
    check(err)
    fmt.Println(string(json))
}

func createBSON(time int64, file string, reqUrl string, callbackUrl string, sName string) []byte {
    _request := Request{
        Url: reqUrl,
        Method: "POST",
        HttpHeaders: map[string]string{
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        Body: map[string]string{
            "name": sName,
        },
    }
    
    input := Input{
        Id: time,
        SuccessHttpResponseCode: 200,
        MaxRetries: 2,
        CallbackWebhookUrl: callbackUrl,
        Request: _request,
    }
    
    inputs := make([]Input, 1)
    inputs[0] = input
    inputArray := InputArray{ Inputs: inputs }
    data, err := bson.Marshal(inputArray)	
    check(err)
    
    return data    
}

func makeTimestamp() int64 {
    return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}

func generateInputBSON(file string, reqUrl string, callbackUrl string, sName string) {
    data := createBSON(makeTimestamp(), file, reqUrl, callbackUrl, sName)
    writeFile(data, file)
    readFile(file)
}

func main() {
    file := os.Args[1]
    reqUrl := os.Args[2]
    callbackUrl := os.Args[3]
    studentName := os.Args[4]
    generateInputBSON(file, reqUrl, callbackUrl, studentName)
}
