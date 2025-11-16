import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
            "@": fileURLToPath(new URL("./src", import.meta.url)),
            "@components": fileURLToPath(new URL("./src/components", import.meta.url)),
            "@data": fileURLToPath(new URL("./src/data", import.meta.url)),
        }
    },
    server: {
        proxy: {
            '/api': {
                target: 'http://192.168.2.189:8080',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '')
            }
        }
    }
});
