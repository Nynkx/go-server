package main

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "os"
	"example.com/m/v2/server"
)

// type Users struct{
// 	Users []User `json:"users"`
// }

// type User struct{
// 	User_name string `json:"user_name"`
// 	Email string     `json:"email"`
// }

// func handlePing(ctx *gin.Context){

// 	jsonFile, err := os.Open("test.json");

// 	if (err != nil){
// 		fmt.Println(err)
// 		ctx.JSON(500,gin.H{
// 			"err":err,
// 		})
// 		return
// 	}

// 	defer jsonFile.Close()

// 	byteVal, _ := ioutil.ReadAll(jsonFile)

// 	var result Users

// 	json.Unmarshal([]byte(byteVal), &result)

// 	ctx.JSON(200, result)
// }

// func init() {
// 	r := gin.Default()

// 	r.GET("/ping", handlePing)
// }

func main() {
	server.Init()
}