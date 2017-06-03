package model

import (
	"database/sql"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/dev/dbcon"
)

type Posts []Post

type Post struct {
	ID        int
	Title     string
	Body      string
	Createdat time.Time
	Updatedat time.Time
}

func (*Posts) All() []Post {
	db := dbcon.Connection()
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		logrus.Warn(err)
		return nil
	}
	defer rows.Close()

	posts := make([]Post, 0)
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Createdat, &post.Updatedat) // order matters
		if err != nil {
			logrus.Warn(err)
			return nil
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		logrus.Warn(err)
		return nil
	}
	return posts
}

func (p *Post) Find(id int) {
	db := dbcon.Connection()
	defer db.Close()

	err := db.QueryRow("SELECT * FROM posts WHERE uid = $1", id).Scan(&p.ID, &p.Title, &p.Body, &p.Createdat, &p.Updatedat)

	switch {
	case err == sql.ErrNoRows:
		logrus.Printf("No post with that ID.")
	case err != nil:
		logrus.Fatal(err)
	}
}
