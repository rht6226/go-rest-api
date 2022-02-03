package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rht6226/go-rest-api/cache"
	"github.com/rht6226/go-rest-api/controller"
	router "github.com/rht6226/go-rest-api/http-router"
	"github.com/rht6226/go-rest-api/repository"
	"github.com/rht6226/go-rest-api/service"
)

var (
	port           string
	httpRouter     router.Router
	repo           repository.PostRepository
	postService    service.PostService
	postCache      cache.PostCache
	postController controller.PostController
)

func init() {
	httpRouter = router.NewMuxRouter()
	repo = repository.NewSQLiteRepository()
	postService = service.NewPostService(repo)
	postCache = cache.NewRedisCache("localhost:6379", 1, 60)
	postController = controller.NewPostController(postService, postCache)

	// port
	assigned := os.Getenv("PORT")
	if len(assigned) != 0 {
		port = assigned
	} else {
		port = "8080"
	}
}

func main() {

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and runing...")
	})

	// register routes
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	err := httpRouter.SERVE(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
}
