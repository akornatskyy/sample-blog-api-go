package mock

import (
	"errors"
	"regexp"
	"strings"

	"github.com/akornatskyy/sample-blog-api-go/posts/domain/post"
	"github.com/akornatskyy/sample-blog-api-go/shared/mock"
)

var (
	reWords = regexp.MustCompile(`\S+`)

	errNotFound = errors.New("not found")
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
			Author: &post.Author{
				FirstName: u.FirstName,
				LastName:  u.LastName,
			},
		})
	}
	return r, nil
}

func (*postRepository) GetPost(slug string) (*post.Post, error) {
	for _, p := range mock.DB.Posts {
		if p.Slug == slug {
			u, ok := mock.DB.UserByID[p.AuthorID]
			if !ok {
				break
			}
			m := &post.Post{
				ID:      p.ID,
				Slug:    p.Slug,
				Title:   p.Title,
				Message: p.Message,
				Created: p.Created,
				Author: &post.Author{
					FirstName: u.FirstName,
					LastName:  u.LastName,
				},
			}
			return m, nil
		}
	}
	return nil, errNotFound
}

func (*postRepository) ListComments(postID, authorID string) ([]*post.Comment, error) {
	var comments []*post.Comment
	for _, c := range mock.DB.Comments {
		if c.PostID != postID || (!c.Moderated && c.AuthorID != authorID) {
			continue
		}
		u, ok := mock.DB.UserByID[c.AuthorID]
		if !ok {
			continue
		}
		comments = append(comments, &post.Comment{
			Author: &post.Author{
				FirstName: u.FirstName,
				LastName:  u.LastName,
			},
			Created:   c.Created,
			Message:   c.Message,
			Moderated: c.Moderated,
		})
	}

	return comments, nil
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
