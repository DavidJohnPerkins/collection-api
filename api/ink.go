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
	Item_id          string `json:"item_id"`
	Map_number       string `json:"map_number"`
	Map_title        string `json:"map_title"`
	Publish_date     string `json:"publish_date"`
	Main_settlements string `json:"main_settlements"`
	Key_value        string `json:"key_value"`
	Map_image        string `json:"map_image"`
	Map_image_rear   string `json:"map_image_rear"`
	Map_image_area   string `json:"map_image_area"`
	Comments         string `json:"comments"`
}

func NewInkResponse(m store.OSMap) osMapResponse {
	return osMapResponse{
		Item_id:          m.Item_id,
		Map_number:       m.Map_number,
		Map_title:        m.Map_title,
		Publish_date:     m.Publish_date,
		Main_settlements: m.Main_settlements,
		Key_value:        m.Key_value,
		Map_image:        m.Map_image,
		Map_image_rear:   m.Map_image_rear,
		Map_image_area:   m.Map_image_area,
		Comments:         m.Comments,
	}
}

func (mr inkResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewInkListResponse(maps []store.OSMap) []render.Renderer {

	list := []render.Renderer{}

	for _, m := range maps {
		mr := NewInkResponse(m)
		list = append(list, mr)
	}
	return list
}

func (s *Server) handleInkList(w http.ResponseWriter, r *http.Request) {

	rangeParam := chi.URLParam(r, "range")
	if rangeParam == "" {
		rangeParam = "LANDRANGER"
	}
	maps, err := s.store.GetOSMapList(r.Context(), rangeParam)
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewInkListResponse(maps))
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
