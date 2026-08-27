/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_LEXFORGE_SERVER_ADDRESS: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
