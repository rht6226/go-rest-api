package cache

import (
	"strconv"
	"testing"

	"github.com/rht6226/go-rest-api/entity"
	"github.com/stretchr/testify/assert"
)

var (
	rCache = NewRedisCache("localhost:6379", 2, 100)
)

const (
	ID    int64  = 72345
	TITLE string = "POST X"
	TEXT  string = "POST X TEST"
)

func TestSetGet(t *testing.T) {
	post := entity.Post{
		Id:    ID,
		Title: TITLE,
		Text:  TEXT,
	}

	rCache.Set(strconv.FormatInt(ID, 10), &post)

	p := rCache.Get(strconv.FormatInt(ID, 10))

	assert.NotNil(t, p)
	assert.Equal(t, ID, p.Id)
	assert.Equal(t, TITLE, p.Title)
	assert.Equal(t, TEXT, p.Text)
}
