package routes

import (
	"devpartner-api/infra/controllers"
	"devpartner-api/infra/database"
	"devpartner-api/infra/database/repository"
	"devpartner-api/infra/middlewares"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func addRoutes(muxR *mux.Router, pool *pgxpool.Pool) {
	NewNotaRouter(controllers.NewNotaController(repository.NewNotaRepository(pool))).Load(muxR)
	muxR.Use(mux.CORSMethodMiddleware(muxR))
}

func NewServer(env map[string]string) *mux.Router {
	mux := mux.NewRouter()

	connect := database.NewConnect(env["DATABASE_URL"])

	addRoutes(mux, connect)

	mux.Use(middlewares.ApplicationTypeSet)

	return mux
}
