.PNONY: openapi-models
# openapiのcomponentsからコードを自動生成します
openapi-models:
	oapi-codegen -config ./openapi/models.config.yaml ./openapi/openapi.yaml

.PNONY: openapi-server
# openapiのpathsからコードを自動生成します
openapi-server:
	oapi-codegen -config ./openapi/server.config.yaml ./openapi/openapi.yaml

.PNONY: openapi-schema
# openapiからコードを自動生成します
openapi-schema: openapi-models openapi-server

.PNONY: up
up:
	docker compose up

.PHONY: prism-mock
## openapiの記述に従ってmockサーバーを起動する
prism-mock:
	$(MAKE) $@ -C frontend/vite-app

.PHONY: prism-mock-d
## openapiの記述に従ってmockサーバーを起動する。レスポンスデータはランダムに生成される。
prism-mock-d:
	$(MAKE) $@ -C frontend/vite-app

.PHONT: prism-mock-local-proxy
## openapiの記述に従ってproxyモードでmockサーバーを起動する
prism-mock-local-proxy:
	$(MAKE) $@ -C frontend/vite-app
