package logic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// response struct
type Response struct {
	Joke string `json:"joke"`
}

func GetJoke() (string, error) {
	fmt.Println("Calling API...")

	// create a http client and create http req
	client := &http.Client{}                                               // client struct
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil) // http request
	if err != nil {
		return "", err
	}

	// add http headers to our request before sending the http request with `response,err :=client.Do(req)`
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// read n the response and then unmarshal it into our response struct
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseObject Response
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return "", err
	}

	return responseObject.Joke, nil
}
