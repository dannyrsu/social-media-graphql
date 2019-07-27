package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dannyrsu/social-media-graphql/models"

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

func (*server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		log.Fatalf("Error creating user: %v\n", err)
		WriteJsonResponse(w, map[string]interface{}{"error": "Invalid Request"})
		return
	}

	response := user.Create()
	WriteJsonResponse(w, response)
}

func (*server) loginHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
		WriteJsonResponse(w, map[string]interface{}{"error": "Login failed"})
	}

	response := models.Login(user.Email, user.Password)
	WriteJsonResponse(w, response)
}

func (*server) createMessageHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	message := &models.Message{}

	err := json.NewDecoder(r.Body).Decode(message)

	if err != nil {
		log.Fatalf("Error creating message: %v\n", err)
		WriteJsonResponse(w, map[string]interface{}{"error": "Error creating message"})
	}

	message.UserID = userID
	response := message.Create()

	WriteJsonResponse(w, response)
}

func (*server) getMessagesByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	messages := models.GetMessagesByEmail(email)

	WriteJsonResponse(w, messages)
}

func (*server) getAllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages := models.GetAllMessages()

	WriteJsonResponse(w, messages)
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
	s.router.Post("/api/user/new", s.createUserHandler)
	s.router.Post("/api/user/login", s.loginHandler)
	s.router.Post("/api/message/new", s.createMessageHandler)
	s.router.Get("/api/user/{email}/message", s.getMessagesByEmailHandler)
	s.router.Get("/api/message", s.getAllMessagesHandler)
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
