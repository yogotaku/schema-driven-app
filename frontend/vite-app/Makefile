.PHONY: prism-mock
## openapiの記述に従ってmockサーバーを起動する
prism-mock:
	prism mock ../../openapi/openapi.yaml

.PHONY: prism-mock-d
## openapiの記述に従ってmockサーバーを起動する。レスポンスデータはランダムに生成される。
prism-mock-d:
	prism mock -d ../../openapi/openapi.yaml

.PHONT: prism-mock-local-proxy
## openapiの記述に従ってproxyモードでmockサーバーを起動する
prism-mock-local-proxy
	prism proxy ../../openapi/openapi.yaml http://localhost:8090
