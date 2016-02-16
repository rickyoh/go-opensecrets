package opensecrets

import (
	"log"
	"errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Apiserver    string
	Key  		 string
	Output   	 string
}

func Connect(Apiserver string, Key string, Output string) *Client {
	return &Client{
		Apiserver:  Apiserver,
		Key:   		Key,
		Output: 	Output,
	}
}

func (c *Client) Query(method string, args []map[string]string) interface{}{

	client := http.Client{}

	url := c.Apiserver
	url = url + "/?method=" + method
	url = url + "&apikey=" + c.Key
	url = url + "&output=" + c.Output

	for _, arg := range args {
		url = url + "&" + arg["key"] + "=" + arg["value"]
	}
	log.Println(url)
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)

	data := map[string]interface{}{}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)

	if(err != nil) {
		log.Println(err)
		errors.New("Problem with opensecrets")
	}
	return data
}