package routes

import (
	"VUTTR-API/infra/controllers"
	"VUTTR-API/infra/database"
	"VUTTR-API/infra/middleware"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func addRoutes(muxR *mux.Router, pool *mongo.Client) {
	NewUserRouter(controllers.NewUserController(database.NewUserRepositoryImp(pool))).Load(muxR)
	NewToolRouter(controllers.NewToolController(database.NewToolRepositoryImp(pool))).Load(muxR)
	NewAuthRouter(controllers.NewAuthController(database.NewUserRepositoryImp(pool))).Load(muxR)
	muxR.Use(mux.CORSMethodMiddleware(muxR))
}

func NewServer(env map[string]string, connect *mongo.Client) *mux.Router {
	mux := mux.NewRouter()

	addRoutes(mux, connect)

	mux.Use(middleware.ApplicationTypeSet)

	return mux
}
