package api

import (
	"dperkins/collection-api/store"
	"log"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type penResponse struct {
	Item_id           int     `json:"item_id"`
	Brand             string  `json:"BRAND"`
	Range             string  `json:"RANGE"`
	Rating            int     `json:"RATING"`
	Key_Value         string  `json:"key_value"`
	Model_name        string  `json:"MODEL_NAME"`
	Body_material     string  `json:"BODY_MATERIAL"`
	Body_colour       string  `json:"BODY_COLOUR"`
	Grip_material     string  `json:"GRIP_MATERIAL"`
	Grip_colour       string  `json:"GRIP_COLOUR"`
	Nib_material      string  `json:"NIB_MATERIAL"`
	Nib_colour        string  `json:"NIB_COLOUR"`
	Nib_size          string  `json:"NIB_SIZE"`
	Purchase_price    float64 `json:"PURCHASE_PRICE"`
	Replacement_price float64 `json:"REPLACEMENT_PRICE"`
	Writing_type      string  `json:"WRITING_TYPE"`
	Cap_material      string  `json:"CAP_MATERIAL"`
	Cap_colour        string  `json:"CAP_COLOUR"`
	Image_1           string  `json:"IMAGE_1"`
	Image_2           string  `json:"IMAGE_2"`
	Comments          string  `json:"comments"`
}

func NewPenResponse(m store.Pen) penResponse {
	return penResponse{
		Item_id:           m.Item_id,
		Brand:             m.Brand,
		Range:             m.Range,
		Rating:            m.Rating,
		Key_Value:         m.Key_Value,
		Model_name:        m.Model_name,
		Body_material:     m.Body_material,
		Body_colour:       m.Body_colour,
		Grip_material:     m.Grip_material,
		Grip_colour:       m.Grip_colour,
		Nib_material:      m.Nib_material,
		Nib_colour:        m.Nib_colour,
		Nib_size:          m.Nib_size,
		Purchase_price:    m.Purchase_price,
		Replacement_price: m.Replacement_price,
		Writing_type:      m.Writing_type,
		Cap_material:      m.Cap_material,
		Cap_colour:        m.Cap_colour,
		Image_1:           m.Image_1,
		Image_2:           m.Image_2,
		Comments:          m.Comments,
	}
}

func (pr penResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewPenListResponse(pens []store.Pen) []render.Renderer {

	list := []render.Renderer{}

	for _, p := range pens {
		pr := NewPenResponse(p)
		list = append(list, pr)
	}
	return list
}

func (s *Server) handlePenList(w http.ResponseWriter, r *http.Request) {

	pens, err := s.store.GetPenList(r.Context())
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewPenListResponse(pens))
}

func (s *Server) handlePenItem(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "item_id")
	id, _ := strconv.Atoi(idParam)

	p, err := s.store.GetPenItem(r.Context(), id)
	if err != nil {
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Render(w, r, NewPenResponse(p))
}
