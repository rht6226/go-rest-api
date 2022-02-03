package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rht6226/go-rest-api/cache"
	"github.com/rht6226/go-rest-api/entity"
	"github.com/rht6226/go-rest-api/errors"
	"github.com/rht6226/go-rest-api/service"
)

// interface
type PostController interface {
	GetPostById(http.ResponseWriter, *http.Request)
	GetPosts(http.ResponseWriter, *http.Request)
	AddPost(http.ResponseWriter, *http.Request)
}

// controller
type postController struct {
	postService service.PostService
	postCache   cache.PostCache
}

// New PostController
func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	return &postController{
		postService: service,
		postCache:   cache,
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

// get post by ID
func (c *postController) GetPostById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	postID := strings.Split(request.URL.Path, "/")[2]

	var post *entity.Post = c.postCache.Get(postID)
	if post == nil {
		post, err := c.postService.FindByID(postID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "No Posts found!"})
			return
		}

		c.postCache.Set(postID, post)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	}
}
