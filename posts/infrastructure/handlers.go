package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/goext/binding"
	"github.com/akornatskyy/goext/httpjson"
	"github.com/akornatskyy/goext/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/posts/usecase/getpost"
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

func GetPostHandler(t httptoken.Token) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		req := &getpost.Request{
			Slug:   p.ByName("slug"),
			Fields: r.URL.Query().Get("fields"),
		}
		if req.Fields != "" {
			t.Authorize(w, r, &req.Principal)
		}

		resp, err := getpost.Process(req)
		if err != nil {
			switch err {
			case getpost.ErrNotFound:
				httpjson.Encode(w, err, http.StatusNotFound)
			default:
				httpjson.Encode(w, err, http.StatusBadRequest)
			}
			return
		}

		httpjson.Encode(w, resp, http.StatusOK)
	}
}

func AddPostCommentHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	}
}
