package main

import (
	"errors"
	"testing"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
)

func Test_processData(t *testing.T) {
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
			guard := monkey.Patch(getDataFromAPI, testCase.GetData)
			defer guard.Unpatch()

			result, err := processData()
			assert.Equal(t, testCase.ExpectedResult, result)
			assert.Equal(t, testCase.ExpectedErr, err)
		})
	}
}
