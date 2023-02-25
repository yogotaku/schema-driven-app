# schema-driven-app

スキーマ駆動開発のサンプルアプリです

```
.
├── app
│   └── src # goのコード
├── frontend # フロントエンドのコード
└── openapi # openapiによるAPI仕様の置き場
```

## 必要ツール群

- Golang
  - バージョンは`.go-version`を参照
- Node.js
- yarn(https://github.com/yarnpkg/yarn)
  - `npm install --global yarn`
- Docker
- make2help(https://github.com/Songmu/make2help)
  - `go install github.com/Songmu/make2help/cmd/make2help@latest`
- oapi-codegen(https://github.com/deepmap/oapi-codegen)
  - `go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest`
- prism(https://github.com/stoplightio/prism)
  - `yarn global add @stoplight/prism-cli`
- spectral(https://github.com/stoplightio/spectral)
  - `yarn global add @stoplight/spectral-cli`
- dredd(https://github.com/apiaryio/dredd)
  - `yarn global add dredd`

## Docker の起動

`make up`コマンドを実行することで docker 上で API サーバが起動します。

## フロントエンドサーバーの起動

`./frontend/vite-app`上で`yarn dev`コマンドを実行してください。
localhost:5173 でフロントエンドサーバーが起動します。
