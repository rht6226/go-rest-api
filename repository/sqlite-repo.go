package repository

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rht6226/go-rest-api/entity"
)

type sqliteRepo struct{}

// Create new sqlite repo with new empty table for post
func NewSQLiteRepository() PostRepository {
	os.Remove("./posts.db")

	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStatement := `
	CREATE table posts (id integer not null primary key, title text, txt text);
	delete FROM posts;
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStatement)
	}

	return &sqliteRepo{}
}

// save post in db
func (*sqliteRepo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	statement, err := tx.Prepare("INSERT INTO posts(id, title, txt) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer statement.Close()

	_, err = statement.Exec(post.Id, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx.Commit()
	return post, nil
}

// Fetch all posts
func (*sqliteRepo) FindAll() ([]entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, txt FROM posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Post

	for rows.Next() {
		var id int64
		var title string
		var text string
		err = rows.Scan(&id, &title, &text)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		post := entity.Post{
			Id:    id,
			Title: title,
			Text:  text,
		}
		posts = append(posts, post)

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return posts, nil
}

// find post with a given ID
func (*sqliteRepo) FindByID(id string) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	row := db.QueryRow("select id, title, txt from posts where id = ?", id)

	var post entity.Post
	if row != nil {
		var id int64
		var title string
		var text string
		err := row.Scan(&id, &title, &text)
		if err != nil {
			return nil, err
		} else {
			post = entity.Post{
				Id:    id,
				Title: title,
				Text:  text,
			}
		}
	}

	return &post, nil
}


// delete from table
func (*sqliteRepo) Delete(post *entity.Post) error {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := tx.Prepare("delete from posts where id = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(post.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	tx.Commit()
	return nil
}
