package post

type Repository interface {
	SearchPosts(q string, limit, offset int) ([]*Post, error)
	GetPost(slug string) (*Post, error)
	ListComments(postID, authorID string) ([]*Comment, error)
}
