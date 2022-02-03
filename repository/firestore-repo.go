package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/rht6226/go-rest-api/entity"
)

const (
	projectId      = "golang-rest-api-e62ce"
	collectionName = "posts"
)

type repo struct{}

// NewFirestoreRepositopy
func NewFirestoreRepositopy() PostRepository {
	return &repo{}
}

// Save
func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create a Firestore client: ", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatal("Failed to add the post to Collection: ", err)
		return nil, err
	}

	return post, nil
}

// FindAll
func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create a Firestore client: ", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil {
			break
		}
		post := entity.Post{
			Id:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// Delete from db : TODO
func (*repo) Delete(*entity.Post) error {
	return nil
}

//FindByID: TODO
func (r *repo) FindByID(id string) (*entity.Post, error) {
	return nil, nil
}
