.PNONY: build
build:
	docker compose build

.PNONY: up
up:
	docker compose up

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

.PNONY: openapi-lint
# openapi.yamlのlintを実行します
openapi-lint:
	spectral lint --ruleset=./openapi/.spectral.yaml ./openapi/openapi.yaml

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

.PHONY: dredd
## dreddを使用してAPIテストを実施する
dredd:
	dredd --config ./openapi/dredd/dredd.yml

.PHONY: dredd-names
## dreddにおける各テストケースの名前を出力する
dredd-names:
	dredd --config ./openapi/dredd/dredd.yml --names
