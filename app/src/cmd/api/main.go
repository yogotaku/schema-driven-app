package main

import (
	"net/http"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yogotaku/schema-driven-app/app/src/infrastructure"
)

func main() {
	// スキーマの定義を取得
	swagger, _ := infrastructure.GetSwagger()

	// スキーマのServerInterfaceを実装した型を取得
	server := infrastructure.NewApiServer()

	r := chi.NewRouter()

	// リクエストがスキーマの定義に合っているかのバリデーション
	r.Use(oapiMiddleware.OapiRequestValidator(swagger))

	// panicが起きた際に復帰する
	r.Use(middleware.Recoverer)

	infrastructure.HandlerFromMux(server, r)

	http.ListenAndServe(":80", r)

}
