package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_struct "main/internal/struct"
	"net/http"
)

func Default(r *http.Request) _struct.Default {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var t _struct.Default
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println(err)
	}
	return t
}
