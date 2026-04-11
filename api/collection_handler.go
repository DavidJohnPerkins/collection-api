package api

import (
	"dperkins/collection-api/store"
	"errors"
	"net/http"
	"strconv"
	"strings"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type osMapResponse struct {
	Id               int    `json:"id"`
	Is_excluded      bool   `json:"is_excluded"`
	Sobriquet        string `json:"sobriquet"`
	Principal_name   string `json:"principal_name"`
	Hotness_quotient int    `json:"hotness_quotient"`
	Nationality      string `json:"nationality"`
	Ranking          string `json:"ranking"`
	Flags            string `json:"flags"`
	TH_url           string `json:"th_url"`
	Movie_count      int    `json:"movie_count"`
}

func NewOSMapResponse(m store.Model) osMapResponse {
	return osMapResponse{
		Id:               m.Id,
		Is_excluded:      m.Is_excluded,
		Sobriquet:        m.Sobriquet,
		Principal_name:   m.Principal_name,
		Hotness_quotient: m.Hotness_quotient,
		Nationality:      m.Nationality,
		Ranking:          m.Ranking,
		Flags:            m.Flags,
		TH_url:           m.TH_url,
		Movie_count:      m.Movie_count,
	}
}

func NewOSMapListResponse(maps []store.OSMap) []render.Renderer {

	list := []render.Renderer{}

	for _, map := range maps {
		//mr := NewOSMapResponse(map)
		//list = append(list, mr)
	}
	return list
}

func (mr osMapResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleOSMapList(w http.ResponseWriter, r *http.Request) {

	termParam := chi.URLParam(r, "term")
	if termParam == "" {
		termParam = "%"
	} else {
		termParam = strings.Replace(termParam, "~", "%", -1)
	}

	models, err := s.store.GetModelList(r.Context(), termParam)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewModelListResponse(models))
}

func (s *Server) handleGetModel(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idParam)

	model, err := s.store.GetModel(r.Context(), id)
	if err != nil {
		var rnfErr *store.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			render.Render(w, r, ErrRecordNotFound)
		} else {
			render.Render(w, r, ErrInternalServerError)
		}
		return
	}

	render.Render(w, r, NewModelExtendedResponse(model))
}

func (s *Server) handleMovieList(w http.ResponseWriter, r *http.Request) {

	movies, err := s.store.GetMovieList(r.Context())
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewMovieListResponse(movies))
}

func (s *Server) handleAttrDescList(w http.ResponseWriter, r *http.Request) {

	termParam := chi.URLParam(r, "attr_abbrev")

	desc, err := s.store.GetAttrDescList(r.Context(), termParam)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewAttrDescListResponse(desc))
}

func (s *Server) handleFlagList(w http.ResponseWriter, r *http.Request) {

	typeParam := chi.URLParam(r, "flag_type")
	desc, err := s.store.GetFlagList(r.Context(), typeParam)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewFlagListResponse(desc))
}
