package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type jsonResult struct {
	A string
	B int
}

func jsonUnmarshal(r io.Reader) ([]jsonResult, error) {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	results := make([]jsonResult, 0, 50)

	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func jsonDecode(r io.Reader) ([]jsonResult, error) {

	results := make([]jsonResult, 0, 50)

	err := json.NewDecoder(r).Decode(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
