package api

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/api/collection", func(r chi.Router) {
		r.Get("/", s.handleOSMapList)
		r.Route("/maps/{range}", func(r chi.Router) {
			r.Get("/", s.handleOSMapList)
		})
		r.Route("/maps/{range}/{item_id}", func(r chi.Router) {
			r.Get("/", s.handleOSMapItem)
		})
		r.Get("/inks", s.handleInkList)
		r.Route("/inks/{item_id}", func(r chi.Router) {
			r.Get("/", s.handleInkItem)
		})
		r.Get("/pens", s.handlePenList)
		r.Route("/pens/{item_id}", func(r chi.Router) {
			r.Get("/", s.handlePenItem)
		})
	})
}
