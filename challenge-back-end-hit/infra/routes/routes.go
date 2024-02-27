package routes

import (
	"challenge-back-end-hit/infra/controllers"
	"challenge-back-end-hit/infra/database"
	"challenge-back-end-hit/infra/middleware"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func addRoutes(muxR *mux.Router, pool *pgxpool.Pool) {
	NewPlanetRouter(controllers.NewPlanetController(pool)).Load(muxR)
	muxR.Use(mux.CORSMethodMiddleware(muxR))
}

func NewServer(env map[string]string) *mux.Router {
	mux := mux.NewRouter()

	connect := database.NewConnect(env["DATABASE_URL"])

	addRoutes(mux, connect)

	mux.Use(middleware.ApplicationTypeSet)

	return mux
}
