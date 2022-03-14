package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	User_name string `json:"user_name"`
	Email     string `json:"email"`
}

func (h Users) GetAll() (*Users, error) {
	jsonFile, err := os.Open("test.json")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var result *Users

	json.Unmarshal([]byte(byteVal), &result)

	return result, nil
}
