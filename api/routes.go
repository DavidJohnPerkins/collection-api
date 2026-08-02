package api

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/api/collection", func(r chi.Router) {
		r.Get("/maplist", s.handleOSMapList)
		r.Get("/map/get", s.handleOSMapItem)

		r.Get("/inklist", s.handleInkList)
		r.Get("/ink/get", s.handleInkItem)

		r.Get("/penlist", s.handlePenList)
		r.Get("/pen/get", s.handlePenItem)

		r.Get("/scorelist", s.handleScoreList)
		r.Get("/score/get", s.handleScoreItem)

		r.Get("/polychromlist", s.handlePolychromList)
		r.Get("/polychrom/get", s.handlePolychromItem)
	})
}
