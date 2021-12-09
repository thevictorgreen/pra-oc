//tap.go
package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
)

func tap(filename string, filepath string, linenumber string, function string) {
	httpposturl := "http://localhost:3000/tap"
	now := time.Now()
	location, terr := time.LoadLocation("EST")
	datime := now.In(location)
	if terr != nil {
		fmt.Println(terr)
	}
	jsonData, _ := json.Marshal(map[string]string{
		"time": datime.String(),
		"filename": filename,
		"filepath": filepath,
		"linenumber": linenumber,
		"function": function,
	})
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
    request.Header.Set("Content-Type", "application/json; charset=UTF-8")
    client := &http.Client{}
    response, error := client.Do(request)
    if error != nil {
        panic(error)
    }
    defer response.Body.Close()
    fmt.Println("response Status:", response.Status)
    fmt.Println("response Headers:", response.Header)
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println("response Body:", string(body))
}