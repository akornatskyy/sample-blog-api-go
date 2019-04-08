package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/goext/binding"
	"github.com/akornatskyy/goext/httpjson"
	"github.com/akornatskyy/sample-blog-api-go/posts/usecase/search"
	"github.com/julienschmidt/httprouter"
)

func ListPostsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &search.Request{}
		if err := binding.Bind(req, r.URL.Query()); err != nil {
			httpjson.Encode(w, err, http.StatusBadRequest)
			return
		}

		resp, err := search.Process(req)
		if err != nil {
			httpjson.Encode(w, err, http.StatusBadRequest)
			return
		}

		httpjson.Encode(w, resp, http.StatusOK)
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
