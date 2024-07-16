package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	// httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
	"github.com/ufguff/service/songs"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	//	router.Get("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	songStore := songs.NewStore(s.db)
	songHandler := songs.NewHandler(songStore)
	songHandler.RegisterRoutes(subrouter)

	c := cors.Default()

	handler := c.Handler(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, handler)
}
