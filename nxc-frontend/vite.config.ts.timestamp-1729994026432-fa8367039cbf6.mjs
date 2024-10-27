// vite.config.ts
import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "file:///Users/kanna/GolandProjects/NxcFull/nxc-frontend/node_modules/vite/dist/node/index.js";
import vue from "file:///Users/kanna/GolandProjects/NxcFull/nxc-frontend/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import AutoImport from "file:///Users/kanna/GolandProjects/NxcFull/nxc-frontend/node_modules/unplugin-auto-import/dist/vite.js";
import Components from "file:///Users/kanna/GolandProjects/NxcFull/nxc-frontend/node_modules/unplugin-vue-components/dist/vite.js";
import { NaiveUiResolver } from "file:///Users/kanna/GolandProjects/NxcFull/nxc-frontend/node_modules/unplugin-vue-components/dist/resolvers.js";
import VueSetupExtend from "file:///Users/kanna/GolandProjects/NxcFull/nxc-frontend/node_modules/vite-plugin-vue-setup-extend/dist/index.mjs";
import svgLoader from "file:///Users/kanna/GolandProjects/NxcFull/nxc-frontend/node_modules/vite-svg-loader/index.js";
var __vite_injected_original_import_meta_url = "file:///Users/kanna/GolandProjects/NxcFull/nxc-frontend/vite.config.ts";
var vite_config_default = defineConfig({
  plugins: [
    vue(),
    svgLoader(),
    VueSetupExtend(),
    // 第二步使用插件
    AutoImport({
      imports: [
        "vue",
        {
          "naive-ui": [
            "useDialog",
            "useMessage",
            "useNotification",
            "useLoadingBar"
          ]
        }
      ]
    }),
    Components({
      resolvers: [NaiveUiResolver()]
    })
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", __vite_injected_original_import_meta_url))
    }
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvVXNlcnMva2FubmEvR29sYW5kUHJvamVjdHMvTnhjRnVsbC9ueGMtZnJvbnRlbmRcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZmlsZW5hbWUgPSBcIi9Vc2Vycy9rYW5uYS9Hb2xhbmRQcm9qZWN0cy9OeGNGdWxsL254Yy1mcm9udGVuZC92aXRlLmNvbmZpZy50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vVXNlcnMva2FubmEvR29sYW5kUHJvamVjdHMvTnhjRnVsbC9ueGMtZnJvbnRlbmQvdml0ZS5jb25maWcudHNcIjtpbXBvcnQge2ZpbGVVUkxUb1BhdGgsIFVSTH0gZnJvbSAnbm9kZTp1cmwnXG5cbmltcG9ydCB7ZGVmaW5lQ29uZmlnfSBmcm9tICd2aXRlJ1xuaW1wb3J0IHZ1ZSBmcm9tICdAdml0ZWpzL3BsdWdpbi12dWUnXG5cbi8vIFx1NUI4OVx1ODhDNU5haXZlVUlcbmltcG9ydCBBdXRvSW1wb3J0IGZyb20gJ3VucGx1Z2luLWF1dG8taW1wb3J0L3ZpdGUnXG5pbXBvcnQgQ29tcG9uZW50cyBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy92aXRlJ1xuaW1wb3J0IHsgTmFpdmVVaVJlc29sdmVyIH0gZnJvbSAndW5wbHVnaW4tdnVlLWNvbXBvbmVudHMvcmVzb2x2ZXJzJ1xuLy8gXHU3QjJDXHU0RTAwXHU2QjY1XHU1QkZDXHU1MTY1XHU2M0QyXHU0RUY2XG5pbXBvcnQgVnVlU2V0dXBFeHRlbmQgZnJvbSAndml0ZS1wbHVnaW4tdnVlLXNldHVwLWV4dGVuZCdcblxuaW1wb3J0IHN2Z0xvYWRlciBmcm9tICd2aXRlLXN2Zy1sb2FkZXInO1xuXG4vLyBodHRwczovL3ZpdGVqcy5kZXYvY29uZmlnL1xuZXhwb3J0IGRlZmF1bHQgZGVmaW5lQ29uZmlnKHtcbiAgICBwbHVnaW5zOiBbXG4gICAgICAgIHZ1ZSgpLFxuICAgICAgICBzdmdMb2FkZXIoKSxcbiAgICAgICAgVnVlU2V0dXBFeHRlbmQoKSwgLy8gXHU3QjJDXHU0RThDXHU2QjY1XHU0RjdGXHU3NTI4XHU2M0QyXHU0RUY2XG4gICAgICAgIEF1dG9JbXBvcnQoe1xuICAgICAgICAgICAgaW1wb3J0czogW1xuICAgICAgICAgICAgICAgICd2dWUnLFxuICAgICAgICAgICAgICAgIHtcbiAgICAgICAgICAgICAgICAgICAgJ25haXZlLXVpJzogW1xuICAgICAgICAgICAgICAgICAgICAgICAgJ3VzZURpYWxvZycsXG4gICAgICAgICAgICAgICAgICAgICAgICAndXNlTWVzc2FnZScsXG4gICAgICAgICAgICAgICAgICAgICAgICAndXNlTm90aWZpY2F0aW9uJyxcbiAgICAgICAgICAgICAgICAgICAgICAgICd1c2VMb2FkaW5nQmFyJ1xuICAgICAgICAgICAgICAgICAgICBdXG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgXVxuICAgICAgICB9KSxcbiAgICAgICAgQ29tcG9uZW50cyh7XG4gICAgICAgICAgICByZXNvbHZlcnM6IFtOYWl2ZVVpUmVzb2x2ZXIoKV1cbiAgICAgICAgfSlcbiAgICBdLFxuICAgIHJlc29sdmU6IHtcbiAgICAgICAgYWxpYXM6IHtcbiAgICAgICAgICAgICdAJzogZmlsZVVSTFRvUGF0aChuZXcgVVJMKCcuL3NyYycsIGltcG9ydC5tZXRhLnVybCkpXG4gICAgICAgIH1cbiAgICB9XG59KVxuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUFrVSxTQUFRLGVBQWUsV0FBVTtBQUVuVyxTQUFRLG9CQUFtQjtBQUMzQixPQUFPLFNBQVM7QUFHaEIsT0FBTyxnQkFBZ0I7QUFDdkIsT0FBTyxnQkFBZ0I7QUFDdkIsU0FBUyx1QkFBdUI7QUFFaEMsT0FBTyxvQkFBb0I7QUFFM0IsT0FBTyxlQUFlO0FBWmtMLElBQU0sMkNBQTJDO0FBZXpQLElBQU8sc0JBQVEsYUFBYTtBQUFBLEVBQ3hCLFNBQVM7QUFBQSxJQUNMLElBQUk7QUFBQSxJQUNKLFVBQVU7QUFBQSxJQUNWLGVBQWU7QUFBQTtBQUFBLElBQ2YsV0FBVztBQUFBLE1BQ1AsU0FBUztBQUFBLFFBQ0w7QUFBQSxRQUNBO0FBQUEsVUFDSSxZQUFZO0FBQUEsWUFDUjtBQUFBLFlBQ0E7QUFBQSxZQUNBO0FBQUEsWUFDQTtBQUFBLFVBQ0o7QUFBQSxRQUNKO0FBQUEsTUFDSjtBQUFBLElBQ0osQ0FBQztBQUFBLElBQ0QsV0FBVztBQUFBLE1BQ1AsV0FBVyxDQUFDLGdCQUFnQixDQUFDO0FBQUEsSUFDakMsQ0FBQztBQUFBLEVBQ0w7QUFBQSxFQUNBLFNBQVM7QUFBQSxJQUNMLE9BQU87QUFBQSxNQUNILEtBQUssY0FBYyxJQUFJLElBQUksU0FBUyx3Q0FBZSxDQUFDO0FBQUEsSUFDeEQ7QUFBQSxFQUNKO0FBQ0osQ0FBQzsiLAogICJuYW1lcyI6IFtdCn0K
