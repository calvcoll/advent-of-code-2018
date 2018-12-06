package utils

import (
	"io/ioutil"
	"strings"
)

func ReadLines(filepath string) ([]string, error) {
	rawData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	data := string(rawData)
	data = strings.TrimSuffix(data, "\n")
	return strings.Split(data, "\n"), nil
}
