package api

import (
	"dperkins/collection-api/store"
	"net/http"

	"github.com/go-chi/render"
)

type filterResponse struct {
	Filter_value string `json:"filter_value"`
}

func NewFilterResponse(m store.FilterValue) filterResponse {
	return filterResponse{
		Filter_value: m.Filter_value,
	}
}

func NewFilterListResponse(filter_value []store.FilterValue) []render.Renderer {

	list := []render.Renderer{}
	for _, fv := range filter_value {
		mr := NewFilterResponse(fv)
		list = append(list, mr)
	}
	return list
}

func (mr filterResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleFilterList(w http.ResponseWriter, r *http.Request) {

	collectionParam := r.URL.Query().Get("collection")
	columnParam := r.URL.Query().Get("dimcol")

	value, err := s.store.GetFilterList(r.Context(), collectionParam, columnParam)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewFilterListResponse(value))
}
