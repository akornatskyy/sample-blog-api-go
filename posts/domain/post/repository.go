package post

type Repository interface {
	SearchPosts(q string, limit, offset int) ([]*Post, error)
}
