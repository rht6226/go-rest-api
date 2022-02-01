package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rht6226/go-rest-api/entity"
	"github.com/rht6226/go-rest-api/repository"
	"github.com/rht6226/go-rest-api/service"
	"github.com/stretchr/testify/assert"
)

var (
	postRepo repository.PostRepository = repository.NewSQLiteRepository()
	postSrvc service.PostService       = service.NewPostService(postRepo)
	postCtrl PostController            = NewPostController(postSrvc)
)

const (
	ID    int64  = 1123
	TITLE string = "Title 1"
	TEXT  string = "Text 1"
)

func TestAddPost(t *testing.T) {
	// create a new HTTP POST req
	jsonData := []byte(fmt.Sprintf(`{"title": "%v", "text": "%v"}`, TITLE, TEXT))
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonData))

	// Assign HTTP Handler Function (controller AddPost)
	handler := http.HandlerFunc(postCtrl.AddPost)

	// Record HTTP Response(httptest)
	response := httptest.NewRecorder()

	// Dispatch the http request
	handler.ServeHTTP(response, req)

	// Add assertions on HTTP Status code and and response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.Id)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	/// cleanup db
	cleanUp(post)
}

func TestGetPost(t *testing.T) {
	setup()

	req, _ := http.NewRequest("GET", "/posts", nil)
	handler := http.HandlerFunc(postCtrl.GetPosts)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	assert.NotNil(t, posts[0].Id)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	cleanUp(posts[0])
}

func setup() {
	post := entity.Post{
		Id:    ID,
		Title: TITLE,
		Text:  TEXT,
	}
	postRepo.Save(&post)
}

func cleanUp(post entity.Post) {
	postRepo.Delete(&post)
}
