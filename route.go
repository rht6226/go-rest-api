package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/rht6226/go-rest-api/entity"
	"github.com/rht6226/go-rest-api/repository"
)

var (
	repo repository.PostRepository = repository.NewFirestoreRepositopy()
)

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}

	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(posts)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post

	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.Id = rand.Int63()

	repo.Save(&post)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
