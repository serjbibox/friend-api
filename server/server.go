package server

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/serjbibox/friend-api/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	default_port     = "8080"
	FRIEND_HTTP_PORT = "PORT"
)

type Server struct {
	port string
}

func New(p string) *Server {
	httpPort := ":"
	if p == "" {
		if env, ok := os.LookupEnv(FRIEND_HTTP_PORT); !ok {
			httpPort += default_port
		} else {
			httpPort += env
		}
	} else {
		httpPort += p
	}
	return &Server{port: httpPort}
}

func (s *Server) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/*", httpSwagger.Handler(httpSwagger.URL("https://friend-service-api.herokuapp.com/swagger/doc.json")))
	// API definition localhost
	//r.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/user", handlers.GetUser)
		r.Get("/user", handlers.GetUser)
		r.Post("/auth/register", handlers.NewUser)
		r.Post("/auth/login", handlers.Login)
	})
	return http.ListenAndServe(s.port, r)
}
