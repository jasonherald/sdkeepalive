package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Setting structure
type Setting struct {
	ID       string `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	SysID    string `json:"sys_id"`
}

func main() {
	jsonFile, err := os.Open("settings.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Loaded settings.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var setting Setting
	json.Unmarshal(byteValue, &setting)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://"+setting.ID+".service-now.com/api/now/table/sys_user/"+setting.SysID, nil)
	req.SetBasicAuth(setting.User, setting.Password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)

	fmt.Println("Got response: " + s)
}
