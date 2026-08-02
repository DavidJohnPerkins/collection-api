package api

import (
	"dperkins/collection-api/store"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type scoreResponse struct {
	Item_id     int    `json:"item_id"`
	Key_Value   string `json:"key_value"`
	Publisher   string `json:"PUBLISHER"`
	Composer    string `json:"COMPOSER"`
	Work_name   string `json:"WORK_NAME"`
	Work_type   string `json:"WORK_TYPE"`
	Score_size  string `json:"SCORE_SIZE"`
	Score_pages int    `json:"SCORE_PAGES"`
	Comments    string `json:"comments"`
}

func NewScoreResponse(m store.Score) scoreResponse {
	return scoreResponse{
		Item_id:     m.Item_id,
		Key_Value:   m.Key_Value,
		Publisher:   m.Publisher,
		Composer:    m.Composer,
		Work_name:   m.Work_name,
		Work_type:   m.Work_type,
		Score_size:  m.Score_size,
		Score_pages: m.Score_pages,
		Comments:    m.Comments,
	}
}

func (ir scoreResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewScoreListResponse(scores []store.Score) []render.Renderer {

	list := []render.Renderer{}

	for _, s := range scores {
		sr := NewScoreResponse(s)
		list = append(list, sr)
	}
	return list
}

func (s *Server) handleScoreList(w http.ResponseWriter, r *http.Request) {

	scores, err := s.store.GetScoreList(r.Context())
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewScoreListResponse(scores))
}

func (s *Server) handleScoreItem(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("item_id")
	id, _ := strconv.Atoi(idParam)
	log.Printf("id: %v", id)
	sc, err := s.store.GetScoreItem(r.Context(), id)
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Render(w, r, NewScoreResponse(sc))
}
