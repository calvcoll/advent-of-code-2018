package utils

import (
	"encoding/json"
	"log"
	"strconv"
)

func ToJsonString(int interface{}) (jsonString string) {
	bytes, e := json.Marshal(int)
	if e != nil {
		log.Fatalln(e)
	}
	return string(bytes)
}

func LogErrorAtoI(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}
