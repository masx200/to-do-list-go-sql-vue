import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig(({ mode, command }) => {
    const isdrop = mode === "production" && command === "build";

    return {
        esbuild: {
            legalComments: "none",
            drop: isdrop ? ["console", "debugger"] : undefined,
        },
        plugins: [
            vue(),
            AutoImport({
                resolvers: [ElementPlusResolver()],
            }),
            Components({
                resolvers: [ElementPlusResolver()],
            }),
        ],
        server: {
            port: 4173,
        },
        preview: {
            port: 4173,
        },
    };
});
