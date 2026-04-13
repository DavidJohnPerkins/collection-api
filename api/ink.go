package api

import (
	"dperkins/collection-api/store"
	"log"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type inkResponse struct {
	Item_id        int    `json:"item_id"`
	Brand          string `json:"BRAND"`
	Range          string `json:"RANGE"`
	Container      string `json:"CONTAINER"`
	Container_size string `json:"CONTAINER_SIZE"`
	Colour         string `json:"COLOUR"`
	Qty_remaining  int    `json:"QTY_REMAINING"`
	Rating         int    `json:"RATING"`
	Ink_name       string `json:"INK_NAME"`
	Key_Value      string `json:"key_value"`
	Image_1        string `json:"IMAGE_1"`
	Comments       string `json:"comments"`
}

func NewInkResponse(m store.Ink) inkResponse {
	return inkResponse{
		Item_id:        m.Item_id,
		Brand:          m.Brand,
		Range:          m.Range,
		Container:      m.Container,
		Container_size: m.Container_size,
		Colour:         m.Colour,
		Qty_remaining:  m.Qty_remaining,
		Rating:         m.Rating,
		Ink_name:       m.Ink_name,
		Key_Value:      m.Key_Value,
		Image_1:        m.Image_1,
		Comments:       m.Comments,
	}
}

func (mr inkResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewInkListResponse(maps []store.Ink) []render.Renderer {

	list := []render.Renderer{}

	for _, m := range maps {
		mr := NewInkResponse(m)
		list = append(list, mr)
	}
	return list
}

func (s *Server) handleInkList(w http.ResponseWriter, r *http.Request) {

	inks, err := s.store.GetInkList(r.Context())
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewInkListResponse(inks))
}

func (s *Server) handleInkItem(w http.ResponseWriter, r *http.Request) {

	rangeParam := chi.URLParam(r, "range")
	if rangeParam == "" {
		rangeParam = "LANDRANGER"
	}
	idParam := chi.URLParam(r, "item_id")
	id, _ := strconv.Atoi(idParam)

	m, err := s.store.GetOSMapItem(r.Context(), rangeParam, id)
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Render(w, r, NewOSMapResponse(m))
}
