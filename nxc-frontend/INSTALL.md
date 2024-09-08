#### 安装NaiveUI

```shell
npm i -D naive-ui
```
#### 安装NaiveUI字体
```shell
npm i -D vfonts
```

#### 按需引入
```ts
// vite.config.ts
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
```

