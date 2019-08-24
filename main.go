package main

import "log"

func main() {
	result, err := processData(getDataFromAPI)
	if err != nil {
		log.Println(err)
	}

	log.Println(result)
}

func processData(getData func() (*Data, error)) (string, error) {
	data, err := getData()
	if err != nil {
		return "", err
	}

	result := "incomplete"
	if data.IsCompleted {
		result = "done :)"
	}

	return result, nil
}
