package main

import "os"
import "io/ioutil"
import "fmt"
import "encoding/base64"
import "net/http"
import "encoding/json"
import "bytes"

type Postdata struct {
	Name     string
	Contents string
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Dude, gives file!")
		os.Exit(1)
	}

	buffer, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	encodeddata := base64.StdEncoding.EncodeToString(buffer)

	postdata := &Postdata{args[1], encodeddata}

	buf, _ := json.Marshal(postdata)
	body := bytes.NewBuffer(buf)
	r, err := http.Post("http://127.0.0.1:8082/test", "application/json", body)
	if err != nil {
		fmt.Println("%s", err)
		panic(err)
	}
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("%s", err)
		panic(err)
	}
	fmt.Println(string(response))

}
