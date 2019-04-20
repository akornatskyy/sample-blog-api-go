package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/goext/binding"
	"github.com/akornatskyy/goext/httpjson"
	"github.com/akornatskyy/goext/httptoken"
	"github.com/akornatskyy/sample-blog-api-go/posts/usecase/addcomment"
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

func AddPostCommentHandler(t httptoken.Token) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var req addcomment.Request
		if err := httpjson.Decode(r, &req, 384); err != nil {
			httpjson.Encode(w, err, http.StatusUnprocessableEntity)
			return
		}
		req.Slug = p.ByName("slug")
		if err := t.Authorize(w, r, &req.Principal); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := addcomment.Process(&req); err != nil {
			switch err {
			case addcomment.ErrForbidden:
				httpjson.Encode(w, err, http.StatusForbidden)
			case addcomment.ErrNotFound:
				httpjson.Encode(w, err, http.StatusNotFound)
			default:
				httpjson.Encode(w, err, http.StatusBadRequest)
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
