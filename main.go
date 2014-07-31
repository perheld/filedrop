package main

import "os"
import "io/ioutil"
import "fmt"
import "encoding/base64"
import "net/http"
import "encoding/json"
import "bytes"

type Postdata struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Contents string `json:"contents"`
}

type ResponseData struct {
	FileName string
}

func main() {
	args := os.Args
	url := "http://somexampleurl/"

	if len(args) < 2 {
		fmt.Println("Dude, gives file!")
		os.Exit(1)
	}

	buffer, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	encodeddata := base64.StdEncoding.EncodeToString(buffer)

	postdata := &Postdata{args[1], "", "data:;base64," + encodeddata}

	buf, _ := json.Marshal(postdata)
	temp := string(buf)
	body := bytes.NewBufferString(temp)

	r, err := http.Post(url+"upload", "application/json", body)
	if err != nil {
		fmt.Println("%s", err)
	}
	response, err := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if err != nil {
		fmt.Println("%s", err)
		panic(err)
	}

	res := &ResponseData{}
	json.Unmarshal([]byte(response), &res)
	fmt.Println(url + "d/" + res.FileName)

}
