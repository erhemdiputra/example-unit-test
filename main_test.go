package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processData(t *testing.T) {
	var getDataFromAPIOriginal = GetDataFromAPI

	defer func() {
		GetDataFromAPI = getDataFromAPIOriginal
	}()

	testCases := []struct {
		Name           string
		GetData        func() (*Data, error)
		ExpectedResult string
		ExpectedErr    error
	}{
		{
			Name: "Get Data return error",
			GetData: func() (*Data, error) {
				return nil, errors.New("error get data")
			},
			ExpectedResult: "",
			ExpectedErr:    errors.New("error get data"),
		},
		{
			Name: "Result is incomplete",
			GetData: func() (*Data, error) {
				return &Data{}, nil
			},
			ExpectedResult: "incomplete",
			ExpectedErr:    nil,
		},
		{
			Name: "Result is incomplete",
			GetData: func() (*Data, error) {
				return &Data{IsCompleted: true}, nil
			},
			ExpectedResult: "done :)",
			ExpectedErr:    nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			GetDataFromAPI = testCase.GetData

			result, err := processData()
			assert.Equal(t, testCase.ExpectedResult, result)
			assert.Equal(t, testCase.ExpectedErr, err)
		})
	}
}
