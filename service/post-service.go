package service

import (
	"errors"
	"math/rand"

	"github.com/rht6226/go-rest-api/entity"
	"github.com/rht6226/go-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct {
	repo repository.PostRepository
}

// Create new PostService
func NewPostService(repo repository.PostRepository) PostService {
	return &service{
		repo: repo,
	}
}

// Validate post
func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}
	return nil
}

// create a new Post
func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()

	return s.repo.Save(post)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}
