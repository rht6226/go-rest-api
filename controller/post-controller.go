package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rht6226/go-rest-api/entity"
	"github.com/rht6226/go-rest-api/errors"
	"github.com/rht6226/go-rest-api/service"
)

// interface
type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

// controller
type postController struct {
	postService service.PostService
}

// New PostController
func NewPostController(service service.PostService) PostController {
	return &postController{
		postService: service,
	}
}

// getPosts
func (c *postController) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := c.postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error fetching the posts"})
		return
	}

	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(posts)
}

// addPost
func (c *postController) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post

	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error decoding the post"})
		return
	}

	err = c.postService.Validate(&post)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	result, err := c.postService.Create(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
