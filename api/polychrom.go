package api

import (
	"dperkins/collection-api/store"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type polychromResponse struct {
	Item_id       int    `json:"item_id"`
	Key_Value     string `json:"key_value"`
	Colour_number string `json:"COLOUR_NUMBER"`
	Colour_name   string `json:"COLOUR_NAME"`
	Length        int    `json:"LENGTH"`
	Colour_group  string `json:"COLOUR_GROUP"`
	Location      string `json:"LOCATION"`
	RGB_red       int    `json:"RGB_RED"`
	RGB_green     int    `json:"RGB_GREEN"`
	RGB_blue      int    `json:"RGB_BLUE"`
	Web_name      string `json:"web_name"`
	Image_1       string `json:"IMAGE_1"`
	Comments      string `json:"comments"`
}

func NewPolychromResponse(p store.Polychrom) polychromResponse {
	return polychromResponse{
		Item_id:       p.Item_id,
		Key_Value:     p.Key_Value,
		Colour_number: p.Colour_number,
		Colour_name:   p.Colour_name,
		Length:        p.Length,
		Colour_group:  p.Colour_group,
		Location:      p.Location,
		RGB_red:       p.RGB_red,
		RGB_green:     p.RGB_green,
		RGB_blue:      p.RGB_blue,
		Web_name:      p.Web_name,
		Image_1:       p.Image_1,
		Comments:      p.Comments,
	}
}

func (ir polychromResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewPolychromListResponse(pencils []store.Polychrom) []render.Renderer {

	list := []render.Renderer{}

	for _, p := range pencils {
		pr := NewPolychromResponse(p)
		list = append(list, pr)
	}
	return list
}

func (s *Server) handlePolychromList(w http.ResponseWriter, r *http.Request) {

	pencils, err := s.store.GetPolychromList(r.Context())
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewPolychromListResponse(pencils))
}

func (s *Server) handlePolychromItem(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("item_id")
	id, _ := strconv.Atoi(idParam)
	log.Printf("id: %v", id)
	pencil, err := s.store.GetPolychromItem(r.Context(), id)
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Render(w, r, NewPolychromResponse(pencil))
}
