package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: 5 * time.Second,
	}

	url = "https://jsonplaceholder.typicode.com/todos/1"
)

type Data struct {
	UserID      int64  `json:"userId"`
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"completed"`
}

func getDataFromAPI() (*Data, error) {
	var data *Data

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("status code not ok")
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
