package service

import (
	"testing"

	"github.com/rht6226/go-rest-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repository
type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}
func (mock *mockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}
func (mock *mockRepository) Delete(*entity.Post) error {
	return nil
}
func (mock *mockRepository) FindByID(id string) (*entity.Post, error) {
	return nil, nil
}

// Case when empty post is supplied
func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

// case when empty title is supplied
func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Id: 1, Title: "", Text: "B"}
	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

// Find all
func TestFindAll(t *testing.T) {
	mockRepo := new(mockRepository)

	var id int64 = 1
	post := entity.Post{Id: id, Title: "A", Text: "B"}

	// setup Expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, id, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

// Create
func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)
	post := entity.Post{Title: "A", Text: "B"}

	// setup Expectation
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	// Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, result.Id)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}
