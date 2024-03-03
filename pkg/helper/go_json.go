package helper

import (
	"encoding/json"
	"fmt"
)

func GoStructToJson(data interface{}) (string, error) {
	jsonMapAsStringFormat, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", jsonMapAsStringFormat)
	return string(jsonMapAsStringFormat), err
}
