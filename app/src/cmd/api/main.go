package main

import (
	"net/http"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/yogotaku/schema-driven-app/app/src/infrastructure"
	"github.com/yogotaku/schema-driven-app/app/src/schema"
)

func main() {
	// スキーマの定義を取得
	swagger, _ := schema.GetSwagger()

	// スキーマのServerInterfaceを実装した型を取得
	server := infrastructure.NewApiServer()

	r := chi.NewRouter()

	// リクエストがスキーマの定義に合っているかのバリデーション
	r.Use(oapiMiddleware.OapiRequestValidator(swagger))

	// panicが起きた際に復帰する
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8100"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}))

	schema.HandlerFromMux(server, r)

	http.ListenAndServe(":80", r)

}
