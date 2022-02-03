package cache

import "github.com/rht6226/go-rest-api/entity"

// interface for caching
type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}

