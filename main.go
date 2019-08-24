package main

import "log"

func main() {
	result, err := processData()
	if err != nil {
		log.Println(err)
	}

	log.Println(result)
}

func processData() (string, error) {
	data, err := GetDataFromAPI()
	if err != nil {
		return "", err
	}

	result := "incomplete"
	if data.IsCompleted {
		result = "done :)"
	}

	return result, nil
}
