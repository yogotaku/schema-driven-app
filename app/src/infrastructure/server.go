package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	// oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
)

// APIサーバーの起動
func RunServer() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORS設定
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8100", "http://localhost:5173", "172.19.0.1:6364"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}))

	// スキーマの定義を取得
	// swagger, _ := schema.GetSwagger()

	// リクエストがスキーマの定義に合っているかのバリデーション
	// r.Use(oapiMiddleware.OapiRequestValidator(swagger))

	// リクエストがスキーマの定義に合っているかのバリデーション
	// r.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, &oapiMiddleware.Options{
	// 	ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(statusCode)
	// 		json.NewEncoder(w).Encode(message)
	// 	},
	// }))

	http.ListenAndServe(":80", r)
}
