package mock

import (
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	user "github.com/akornatskyy/sample-blog-api-go/membership/infrastructure/mock"
	"github.com/akornatskyy/sample-blog-api-go/posts/domain/post"
	"github.com/akornatskyy/sample-blog-api-go/shared/iojson"
)

type postRepository struct {
}

type postInfo struct {
	ID       string    `json:"id"`
	Created  time.Time `json:"created_on"`
	AuthorID string    `json:"author_id"`
	Slug     string    `json:"slug"`
	Title    string    `json:"title"`
	Message  string    `json:"message"`
}

type commentInfo struct {
	ID        string    `json:"id"`
	Created   time.Time `json:"created_on"`
	AuthorID  string    `json:"author_id"`
	Moderated bool      `json:"moderated"`
	PostID    string    `json:"post_id"`
	Message   string    `json:"message"`
}

type data struct {
	Posts    []*postInfo    `json:"posts"`
	Comments []*commentInfo `json:"comments"`
}

var (
	samples = loadPostSamples("post-samples.json")

	errNotFound = errors.New("not found")
)

func NewPostRepository() post.Repository {
	return &postRepository{}
}

func (*postRepository) SearchPosts(q string, limit, offset int) ([]*post.Post, error) {
	posts := filter(q, limit, offset)
	r := make([]*post.Post, 0, limit)
	for _, p := range posts {
		u, err := user.FindUserByID(p.AuthorID)
		if err != nil {
			continue
		}
		r = append(r, &post.Post{
			Slug:    p.Slug,
			Title:   p.Title,
			Message: truncateWords(p.Message, 40),
			Created: p.Created,
			Author: post.Author{
				FirstName: u.FirstName,
				LastName:  u.LastName,
			},
		})
	}
	return r, nil
}

func filter(q string, limit, offset int) []*postInfo {
	if q == "" {
		n := len(samples.Posts)
		if offset >= n {
			return nil
		}
		e := offset + limit
		if e > n {
			e = n
		}
		return samples.Posts[offset:e]
	}
	posts := make([]*postInfo, 0, limit)
	q = strings.ToLower(q)
	n := offset + limit
	i := 0
	for _, p := range samples.Posts {
		if strings.Contains(strings.ToLower(p.Title), q) {
			if i >= offset {
				posts = append(posts, p)
			}
			i++
			if i == n {
				break
			}
		}
	}
	return posts
}

var reWords = regexp.MustCompile(`\S+`)

func truncateWords(s string, c int) string {
	w := reWords.FindAllString(strings.ReplaceAll(s, "\\n", " "), c)
	if len(w) == c {
		w = append(w, "...")
	}
	return strings.Join(w, " ")
}

func loadPostSamples(filename string) data {
	var d data
	if err := iojson.ReadFile(filename, &d); err != nil {
		panic(err)
	}

	log.Printf("loaded %d posts", len(d.Posts))
	log.Printf("loaded %d comments", len(d.Comments))
	return d
}
