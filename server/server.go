package server

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type server struct {
	router *chi.Mux
}

func (*server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("GraphQL server"))
}

func (s *server) middleware() {
	s.router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
		JwtValidate,
	)
}

func (s *server) routes() {
	s.router.Get("/", s.defaultHandler)
}

func InitializeServer() http.Handler {
	server := &server{
		router: chi.NewRouter(),
	}

	server.middleware()
	server.routes()

	handler := cors.Default().Handler(server.router)

	return handler
}
