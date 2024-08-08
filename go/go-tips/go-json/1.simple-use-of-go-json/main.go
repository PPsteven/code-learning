package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//1. struct to string
//2. string to struct
//3. load json from file

type User struct {
	FirstName string `json:"first_name"` // key will be "first_name"
	BirthYear int    `json:"birth_year"` // key will be "birth_year"
	Email     string // key will be "Email"
}

//1. struct to string
func structToString(user *User) string {
	userByte, _ := json.Marshal(user)
	return string(userByte)
}

//2. string to struct
func stringToStruct(s string) (user *User) {
	_ = json.Unmarshal([]byte(s), &user)
	return user
}

//3. load json from file
func readJsonFromFile(path string) (user *User) {
	userByte, _ := ioutil.ReadFile(path)
	_ = json.Unmarshal(userByte, &user)
	return user
}

func main() {
	//1. struct to string
	userString := structToString(&User{"ppsteven", 2021, "ppsteven@outlook.com"})
	fmt.Printf("%v\n", userString)
	// {"first_name":"ppsteven","birth_year":2021,"Email":"ppsteven@outlook.com"}

	//2. string to struct
	user := stringToStruct(userString)
	fmt.Printf("%#v\n", user)
	// &main.User{FirstName:"ppsteven", BirthYear:2021, Email:"ppsteven@outlook.com"}

	//3. load json from file
	tempFile, _ := ioutil.TempFile("/tmp", "temp-*.json")
	defer os.Remove(tempFile.Name())
	_, _ = tempFile.WriteString(userString)

	user = readJsonFromFile(tempFile.Name())
	fmt.Printf("%#v\n", user)
	// &main.User{FirstName:"ppsteven", BirthYear:2021, Email:"ppsteven@outlook.com"}
}