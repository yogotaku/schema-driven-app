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



