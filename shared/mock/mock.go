package mock

import (
	"log"
	"time"

	"github.com/akornatskyy/sample-blog-api-go/public/domain/quote"
	"github.com/akornatskyy/sample-blog-api-go/shared/iojson"
)

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	IsLocked     bool   `json:"is_locked"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type Post struct {
	ID       string    `json:"id"`
	Created  time.Time `json:"created_on"`
	AuthorID string    `json:"author_id"`
	Slug     string    `json:"slug"`
	Title    string    `json:"title"`
	Message  string    `json:"message"`
}

type Comment struct {
	ID        string    `json:"id"`
	Created   time.Time `json:"created_on"`
	AuthorID  string    `json:"author_id"`
	Moderated bool      `json:"moderated"`
	PostID    string    `json:"post_id"`
	Message   string    `json:"message"`
}

var (
	DB struct {
		Users    []*User        `json:"users"`
		Posts    []*Post        `json:"posts"`
		Comments []*Comment     `json:"comments"`
		Quotes   []*quote.Quote `json:"quotes"`

		UserByID map[string]*User
	}
)

func Load(filename string) {
	if err := iojson.ReadFile(filename, &DB); err != nil {
		panic(err)
	}

	DB.UserByID = make(map[string]*User, len(DB.Users))
	for _, u := range DB.Users {
		DB.UserByID[u.ID] = u
	}

	log.Printf("loaded: %d users, %d posts, %d comments, %d quotes",
		len(DB.Users), len(DB.Posts), len(DB.Comments), len(DB.Quotes))
}
