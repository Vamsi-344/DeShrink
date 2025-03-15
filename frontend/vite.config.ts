import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
import react from '@vitejs/plugin-react'
import { TanStackRouterVite } from '@tanstack/router-plugin/vite'
// https://vite.dev/config/
export default defineConfig({
  server: {
    proxy: {
    '/api': {
      target: 'http://localhost:8888',
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/api/, ''),
    },
  }},
  plugins: [TanStackRouterVite({ target: 'react', autoCodeSplitting: true }), react(), tailwindcss()],
})
