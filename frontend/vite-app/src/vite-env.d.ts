/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_API_ENDPOINT_URL: string;
  readonly VITE_PRISM_MOCK_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
