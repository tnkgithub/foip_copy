import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"

export default defineConfig({
  // TODO: serve時の処理も書いておきたい(具体的にはbackendのURL)
  plugins: [react()]
})
