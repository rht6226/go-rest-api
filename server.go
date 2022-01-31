package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rht6226/go-rest-api/controller"
	router "github.com/rht6226/go-rest-api/http-router"
	"github.com/rht6226/go-rest-api/repository"
	"github.com/rht6226/go-rest-api/service"
)

var (
	httpRouter     router.Router
	repo           repository.PostRepository
	postService    service.PostService
	postController controller.PostController
)

const (
	port string = ":8080"
)

func init() {
	httpRouter = router.NewMuxRouter()
	repo = repository.NewFirestoreRepositopy()
	postService = service.NewPostService(repo)
	postController = controller.NewPostController(postService)
}

func main() {

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and runing...")
	})

	// register routes
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	err := httpRouter.SERVE(port)
	if err != nil {
		log.Fatal(err)
	}
}
