package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/sample-blog-api-go/posts/usecase/search"
	"github.com/akornatskyy/sample-blog-api-go/shared/binding"
	"github.com/akornatskyy/sample-blog-api-go/shared/httpjson"
	"github.com/julienschmidt/httprouter"
)

func ListPostsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &search.Request{}
		if err := binding.Bind(req, r.URL.Query()); err != nil {
			httpjson.Encode(w, err, http.StatusBadRequest)
			return
		}
	}
}

func GetPostHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	}
}

func ListPostCommentsHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	}
}
