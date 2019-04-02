package mock

import (
	"regexp"
	"strings"

	"github.com/akornatskyy/sample-blog-api-go/shared/mock"

	"github.com/akornatskyy/sample-blog-api-go/posts/domain/post"
)

var (
	reWords = regexp.MustCompile(`\S+`)
)

type postRepository struct {
}

func NewPostRepository() post.Repository {
	return &postRepository{}
}

func (*postRepository) SearchPosts(q string, limit, offset int) ([]*post.Post, error) {
	items := filter(q, limit, offset)
	r := make([]*post.Post, 0, limit)
	for _, p := range items {
		u, ok := mock.DB.UserByID[p.AuthorID]
		if !ok {
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

func filter(q string, limit, offset int) []*mock.Post {
	if q == "" {
		n := len(mock.DB.Posts)
		if offset >= n {
			return nil
		}
		e := offset + limit
		if e > n {
			e = n
		}
		return mock.DB.Posts[offset:e]
	}
	items := make([]*mock.Post, 0, limit)
	q = strings.ToLower(q)
	n := offset + limit
	i := 0
	for _, p := range mock.DB.Posts {
		if strings.Contains(strings.ToLower(p.Title), q) {
			if i >= offset {
				items = append(items, p)
			}
			i++
			if i == n {
				break
			}
		}
	}
	return items
}

func truncateWords(s string, c int) string {
	w := reWords.FindAllString(strings.Replace(s, "\\n", " ", -1), c)
	if len(w) == c {
		w = append(w, "...")
	}
	return strings.Join(w, " ")
}
