package post

type Repository interface {
	SearchPosts(q string, limit, offset int) ([]*Post, error)
	GetPost(slug string) (*Post, error)
	GetPostId(slug string) (string, error)
	ListComments(postID, authorID string) ([]*Comment, error)
	CountCommentsAwaitingModeration(authorID string, limit int) (int, error)
	AddPostComment(postID, authorID, message string) error
}
