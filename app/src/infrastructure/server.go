package infrastructure

import (
	"encoding/json"
	"net/http"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/yogotaku/schema-driven-app/app/src/interface/controllers"
	"github.com/yogotaku/schema-driven-app/app/src/schema"
)

type ApiServer struct {
	*controllers.UserController
	*controllers.PetController
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		UserController: controllers.NewUserController(),
		PetController:  controllers.NewPetController(),
	}
}

func RunServer() {
	// スキーマの定義を取得
	swagger, _ := schema.GetSwagger()

	// スキーマのServerInterfaceを実装した型を取得
	server := NewApiServer()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// panicが起きた際に復帰する
	r.Use(middleware.Recoverer)

	// CORS設定
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8100", "http://localhost:5173", "172.19.0.1:6364"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}))

	// リクエストがスキーマの定義に合っているかのバリデーション
	r.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, &oapiMiddleware.Options{
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(message)
		},
	}))

	schema.HandlerFromMux(server, r)

	http.ListenAndServe(":80", r)
}
