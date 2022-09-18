package main

import (
	"bytes"
	"consume-rest-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	getPosts()
	savePosts()
}

func getPosts() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)
	}

	var responseObject []models.Posts
	json.Unmarshal(responseData, &responseObject)

	//fmt.Println(responseObject)
}

func savePosts() {
	post := models.Posts{
		UserId: 1,
		Body:   "Body",
		Title:  "title",
	}

	json_data, err := json.Marshal(post)
	if err != nil {
		fmt.Print(err)
	}
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(response.StatusCode)
}
