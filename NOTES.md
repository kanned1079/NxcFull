# Vue3笔记

## 创建项目

### 基于Vite创建

- 终端运行 `npm create vue@latest`
- 加入TypeScript支持

#### 注意

- `Vite`项目中 `index.html`是入口文件 在项目的最外层
- 加载`index.html`后 Vite解析`<script type=""module" src="xxx">`指向的JavaScripts
- Vue3通过`createApp()`函数创建一个简单的实例

## Vue3核心语法

### 选项式API和组合式API

- `Vue2`的API设计是`选项式(Options)`风格
- `Vue3`的API设计是`组合式(Composition)`风格

#### 选项式API的弊端

`Options`类型的API 数据、方法、计算属性等 是分散在`data` `methods` `computed`中的 如果想新增或者修改一个需求 就需要分别修改
不利于维护

```html

<script lang="ts">
    export default {
        name: 'Person',
        data() {
            return {...}
        },
        methods: {...},
        computed: {...},
        watch: {...}
    }
</script>

<template>
    <div class="person">
        <h2>姓名： {{ name }}</h2>
        <h2>年龄： {{ age }}</h2>
        <button @click="modifyName">修改名字</button>
        <button @click="modifyAge">修改年龄</button>
        <button @click="showTel">显示联系方式</button>
    </div>
</template>
```

### 拉开序幕的 `setup`

`setup`是Vue3中一个新的配置项 值是一个函数 它是`Composition API`"表演的舞台" 组件中所用到的 数据 方法 计算属性 监视
均配置在`setup`中

#### 特点如下

- `setup`函数返回对西那个中的内容 可直接在模版中使用
- `setup`中访问`this`是`undefined`
- `setup`函数会在`beforeCreate`中调用 它是领先所有hooks执行的

#### `setup`语法糖

- `setup`有一个语法糖 可以让我们把`setup`独立出去
    ```html
    <script lang="ts" setup>
        console.log(this);  // undefined
        let x = 99;
    </script>
    ```
- 如果使用上述写法 还需要写一个不带`setup`的标签来设置name 因此可以借助一个插件来实现简化

##### 使用`vite-plugin-vue-setup-extend`插件

1. 安装插件 `yarn add vite-plugin-vue-setup-extend -D`
2. 编辑 `vite.config.ts` 加入两行内容
    ```typescript
    import {fileURLToPath} from 'node:url'
    
    import {defineConfig} from 'vite'
    import vue from '@vitejs/plugin-vue'
    // 第一步导入插件
    import VueSetupExtend from 'vite-plugin-vue-setup-extend'
    
    // https://vitejs.dev/config/
    export default defineConfig({
        plugins: [
            vue(),
            VueSetupExtend(), // 第二步使用插件
        ],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            }
        }
    })
    ```
3. 在标签中加入`name`属性即可
    ```html
    <script lang="ts" setup name="Person234">
     ... 
    </script>
    ```

### ref创建: 基本类型的响应式数据

- **作用：** 定义响应式变量(基本/对象都可)
- **语法：** `let <变量名> = ref(<初始值>);`
- **返回值：** 一个`RefImpl`实例对象 对象的`value`属性是响应式的
- **注意点：**
    - `JS`中操作数据需要`<变量>.value` 但是模版中不需要`.value`
    - 对于`let name = ref('张三');`来说 `name`不是响应式的 `name.value`是响应式的

### reactive创建: 对象类型的响应式数据

- **作用：** 定义响应式变量(只能用在对象中)
- **语法：** `let <变量名> = reactive({<源对象>});`
- **返回值：** 一个`Proxy`实例对象 简称`响应式对象`
- **注意点：**
    - `reactive`中定义的响应式数据是"深层次"的

### `ref`对比`reactive`

#### 宏观角度

> 1. `ref`用来定义 **基本数据类型**、**对象类型数据**
> 2. `reactive`用来定义 **对象类型数据**

#### 区别

> 1. `ref`创建的变量必须使用`.value`(可以使用volar插件自动添加.value)
> 2. `reactive`重新分配一个对象时 会**失去**响应式 (可以使用`Object.assign`去整体替换)

```typescript
// reactive的情况
car = {brand: '奥拓', price: 12}  // 这就叫重新分配一个对象 会因此失去响应式
Object.assign(car, {brand: '奥拓', price: 12})  // 可以使用Object.assign
// ref的情况 可以直接使用
car.value = {brand: '奥拓', price: 12}
```

#### 使用原则

> 1. 若需要一个基本类型的响应式数据 必须使用`ref`
> 2. 若需要一个响应式对象 层级不深 `ref`和`reactive`都可以
> 3. 若需要一个响应式对象 且层级较深 推荐使用`reactive`

### `toRefs`和`toRef`

#### toRefs函数

- 可以将一个响应式对象转换为一个普通的对象 该对象的每个属性都是独立的`ref`对象
- 返回的对象可以进行解构 每个属性都可以像普通的`ref`对象一样访问和修改，而且会保持响应式的关联
- 使用场景主要是在将响应式对象作为属性传递给子组件时 确保子组件可以正确地访问和更新这些属性
    ```typescript
    let {name, age} = toRefs(person)    // 此时name age都是响应式的ref
    ```

#### toRef函数

- 可以将一个响应式对象的属性转换为一个独立的`ref`对象
- 返回的是一个指向源对象属性的`ref`引用 任何对该引用的修改都会同步到源对象属性上。
- 使用`toRef`时需要传入源对象和属性名作为参数 `toRef(<对象名>, '<对象中的属性名>')`
    ```typescript
    let nl = toRef(person, 'age') 
    ```

### computed计算属性

- 基本语法 `let <计算属性名> = computed(() => { ... return ... })`
- 计算属性有缓存 多次相同值调用只计算一次

#### 定义一个计算属性 只读

```typescript
let fullName = computed(() => {
    console.log('computed');
    return firstName.value.slice(0, 1).toUpperCase() + firstName.value.slice(1) + '-' + lastName.value;
})
```

#### 定义一个计算属性 可读可写

```typescript
// 计算属性
let fullName = computed({
    get() {
        return firstName.value.slice(0, 1).toUpperCase() + firstName.value.slice(1) + '-' + lastName.value;
    },
    set(val) {
        let [str1, str2] = val.split('-');
        firstName.value = str1;
        lastName.value = str2;
    }
})

// 由点击事件触发
let changeFullName = () => {
    fullName.value = 'li-si'
}
```

### watch

- **作用：** 监视数据的变化 (和`Vue2`中的*作用*是一致的)
- **特点：** `Vue3`中的`watch`只能监视以下**四种数据**：
  > 1. `ref`定义的数据
  > 2. `reactive`定义的数据
  > 3. 函数返回一个值(一个`getter()`函数)
  > 4. 一个包含上述内容的数组

#### 语法

- **有三个参数：** `watch(<被监视的数据>, (newVal, oldVal) => {<回调函数>}), {<配置对象>}`

#### 一般遇到的几个情况

##### 情况1

- 监视`ref`定义的**基本类型**数据 直接写数据名即可 不需要`.value`
  ```typescript
  // 数据
  let sum = ref(0)
  // 监视
  const stopWatch = watch(sum, (newValue, oldValue) => {
    console.log('sum changed', newValue, oldValue)
    if (newValue >= 10)
      stopWatch()   // 解除监视
  })
  ```

##### 情况2

- 监视`ref`定义的**对象类型**数据 监视的是对象的**地址值**
- 若想监视其内部数据 要手动开启深度监视
  > 注意
  > - 若修改的`ref`定义的对象的属性 `newValue`和`oldValue`都是新的值
  > - 若修改整个`ref`对象 `newValue`是新值 `oldValue`是旧值 因为不是同一个对象了
  ```typescript
    // 数据
    let person = ref({name: '张三', age: 18});
    let changeName = () => person.value.name += '~';
    let changeAge = () => person.value.age += 1;
    let changePerson = () => person.value = {name: '李四', age: 29};
    // 监视
    watch(person, (newValue, oldValue) => {
      console.log('person变化了', newValue, oldValue)  // 默认监视的是对象的地址值 需要开启深度监视
    }, {
      deep: true, // 需要开启深度监视
      // immediate: true, // 立即执行
    })
  ```

##### 情况3

- 监视`reactive`定义的**对象类型**数据
- 默认开启了深度监视 而且不可关闭
  ```typescript
  // 数据
  let obj = reactive({a: {b: {c: 666,}}})
  // 方法
  let changeVal = () => obj.a.b.c = 999
  // 监视
  watch(obj, (newVal, oldVal) => {
    console.log('obj变化了', newVal, oldVal)
  })
  ```

##### 情况4

- 监视`ref`或`reactive`定义的**对象类型**数据中的**某个属性**有以下注意点
    1. 若该属性**不是** *对象类型* 需要写成函数形式
    2. 若该属性值**依然**是 *对象类型* 可直接编写 也可以写成函数 建议使用函数
- *ps:* 监视的要是对象中的属性 那么最好写函数式
- *注意点：* 若是对象监视的式地址值 需要关注对象内部 需要手动开启深度监视

  ```typescript
  // 只监视对象中的一个属性 该属性是基本类型 需要写成函数式
  watch(() => person.name, (newValue, oldValue) => {
      console.log('person.name变化了', newValue, oldValue)
  })
  // 使用函数式 建议的写法 此时监视的就是对象的地址值了
  watch(() => person.cars, (newValue, oldValue) => {
      console.log('person.cars变化了 函数式', newValue, oldValue)
  }, { deep: true, })
  ```

##### 情况5

- 监视指定的多个属性
  ```typescript
  watch([() => person.name, () => person.cars.c1], (newValue, oldValue) => {
    console.log('person.name或person.cars.c1变化了', newValue, oldValue)
  }, {deep: true,})
  ```

### watchEffect

- 立即运行一个函数 同时响应式地之宗其以来 并在依赖更改时重新执行该函数
- `watch`对比`watchEffect`:
    - 都能监听响应式数据的变化 不同的是监听数据变化的方式不同
    - `watch` 要明确指出监视的数据
    - `watchEffect` 不用明确指出监视的数据 (函数中用到那些属性 那就监视哪些属性)

  ```typescript
  // 使用watch实现
  watch([temp, height], (newValue) => {
    console.log(newValue)
    let [newTemp, newHeight] = newValue;
    console.log(newTemp, newHeight)
    if (newTemp >= 60 || newHeight >= 80)
      console.log('sent req')
  })
  // 使用watchEffect实现
  watchEffect(() => {
    console.log('@'); // 会立即运行一次
    if (temp.value >= 60 || height.value >= 80)
      console.log('sent req')
  })
  ```

### 标签的`ref`标识

- **作用：** 用于注册模版引用
    - 用在普通`DOM`标签上 获取的是`DOM`节点
    - 用在组件标签上 获取的是组件的实例对象

  ```typescript
  let title2 = ref() // 存储ref标记的内容
  // 每个组件间的title2相互独立
  ```

  ```html
  <!--这是用在DOM标签上-->
  <h2 ref="title2">上海</h2>
  <!--这是用在组件实例上-->
  <Person ref="ren"></Person>
  ```

    - 父组件可以获取子组件导出的数据 使用`defineExpose`

  ```typescript
  // 引用defineExpose
  import {ref, defineExpose} from 'vue'
  // ...
  // 在子组件最后进行导出
  defineExpose({a, b, c})
  ```

### TS类型限制`interface`

#### 定义`@/type/index.ts`

```typescript
// 定义一个接口 用于限制Person对象的具体属性
export interface PersonInter {
    id: string
    name: string
    age: number
}

// 一个自定义类型
export type Persons = Array<PersonInter>
```

#### 使用`@/components/Person.vue`

```typescript
// 导入
import {type PersonInter, type Persons} from '@/types'
// 使用
// 单个属性
let person: PersonInter = {id: '0001', name: '张三', age: 60}
// 数组类型
let personList: Array<PersonInter> = [
    // 必须要符合接口规范
    {id: '0001', name: '张三', age: 60},
]
// 另一种写法
let personsList: Persons = [
    {id: '0001', name: '张三', age: 60},
]
```

### 父传递数据 `defineProps`和`withDefaults`的使用

- **作用：** 使用`defineProps`来定义父组件传递的`props`

#### 父组件中设置数据

```typescript
import {type Persons} from "@/types"; // 数据类型限制
// 推荐写法
let personList = reactive<Persons>([
    {id: '0001', name: '张三', age: 18},
    {id: '0002', name: '李四', age: 19},
    {id: '0003', name: '王五', age: 20, x: 999},
])
```

```html

<Person v-bind:list="personList"></Person>
```

#### 在子组件中接收

```typescript
// 导入限制
import {type Persons} from "@/types";
import {withDefaults} from "vue";
// 只接收a
// 但是这样的接收可能会引起类型不匹配问题
// defineProps(['a'])

// 接收a 同时将a保存
// let x = defineProps(['a', 'list'])

// 接收a + 限制类型
// defineProps<{list:Persons}>()

// 接收list 限制类型 限制必要性 指定默认值
withDefaults(defineProps<{ list?: Persons }>(), {
    list: () => { // 默认值需要以函数形式返回
        return [
            {id: '10002', name: 'kanna', age: 21},
            {id: '10003', name: 'kinggyo', age: 16},
            // ...
        ]
    }
})
```

### 组件的生命周期

- **概念：** `Vue`组件实例在创建时要经历一系列的初始化步骤在此过程中`vue`会在合适的时机调用特定的函数
  从而让开发者有机会在特定阶段运行自己的代码 这些特定的函数统称为: 生命周期钩子
- **规律：** 整体分为四个阶段 **创建** **挂载** **更新** **销毁**

#### `Vue2`的生命周期

- 创建`beforeCreate` `created`
- 挂载 `beforeMount` `mounted`
- 更新 `beforeUpdate` `updated`
- 销毁 `beforeDestrory` `destroyed`

#### `Vue3`的生命周期

- 创建阶段 `setup`
- 挂载阶段 `onBeforeMount` `onMounted`
- 更新阶段 `onBeforeUpdate` `onUpdated`
- 卸载阶段 `onBeforeUnmount` `onUnmounted`

#### 常用的钩子

- `onMounted` `onUpdated` `onBeforeUnmount`

#### 示例

```typescript
// 创建 setup中就是创建
console.log('创建');
// 挂载
onBeforeMount(() => console.log('挂载前'))
onMounted(() => console.log('Person挂载完毕'))
// 更新
onBeforeUpdate(() => console.log('更新前'))
onUpdated(() => console.log('更新完毕'))
// 卸载
onBeforeUnmount(() => console.log('卸载前'))
onUnmounted(() => console.log('卸载完毕'))
```

### 自定义Hooks

- 将原本于组件中的dog处理逻辑拆分到`@/hooks/useDog.ts`中

  ```typescript
  import {reactive, onMounted} from "vue";
  import axios from "axios";
  export default () => {
      // 数据
      let dogList = reactive([
          'https://images.dog.ceo/breeds/pembroke/n02113023_12248.jpg'
      ])
      // 方法
      let getDog = async () => {
          try {
              let result = await axios.get('https://dog.ceo/api/breed/pembroke/images/random');
              dogList.push(result.data.message)
          } catch (error) {
              alert(error)
          }
      }
      // 钩子
      onMounted(() => {
          getDog()
      })
      // 向外部提供东西
      return {dogList, getDog}
  }
  ```

- 将原本于组件中的sum处理逻辑拆分到`@/hooks/useSum.ts`中

  ```typescript
  import {onMounted, ref, computed} from "vue";
  export default () => {
      // 数据
      let sum = ref(0)
      // 方法
      let add = () => sum.value += 1;                                          e += 1;
      // 钩子
      onMounted(() => add())
      // 计算属性
      let bigSum = computed(() => sum.value * 10)
      // 向外部提供东西
      return {sum, add, bigSum}
  }
  ```

- 在组件中导入并使用
  ```typescript                                                                                             
  import useDog from "@/hooks/useDog";
  import useSum from "@/hooks/useSum";
  // 解构
  const {sum, add, bigSum} = useSum()
  const {dogList, getDog} = useDog()
  ```

## 路由

### 概念

- 路由就是一组`key` - `value`的对应关系
- 多个路由 需要经过**路由器** 的管理

#### 创建一个路由

- 创建路由文件`@/components/router/index.ts`

  ```typescript
  // 创建一个路由并暴露
  import {createRouter, createWebHistory} from "vue-router";
  // 引入要呈现的组件
  import Home from '@/components/Home.vue'
  import News from '@/components/News.vue'
  import About from '@/components/About.vue'
  // 创建路由器
  const router = createRouter({
      // 配置对象
      history: createWebHistory(),    // 指定路由时候一定要指定路由的工作模式
      routes: [   // 一个个的路由规则
          {
              path: '/home',
              component: Home,
          },
          {
              path: '/news',
              component: News,
          },
          {
              path: '/about',
              component: About,
          },
      ]
  })
  // 导出
  export default router
  ```

- 在`main.ts`中进行引入

  ```typescript
  import { createApp } from "vue";
  import App from './App.vue'
  
  // 引入路由组件
  import router from '@/router'
  // 创建一个应用
  const app = createApp(App)
  // 使用路由组件
  app.use(router)
  // 在最后进行挂载
  app.mount('#app')
  ```

- 在首页使用

  ```typescript
  // 引入RouterView和RouterLink
  import { RouterView, RouterLink } from "vue-router";
  ```

  ```html
  这里是替换了<a>属性
  <RouterLink to="home" active-class="active">首页</RouterLink>
  这里是路由的出口
  <RouterView></RouterView>
  ```

#### 两个注意点

1. 路由组件通常存放在`pages`或`views`文件夹 一般组件通常放置在`components`中
2. 通过点击导航 视觉上消失了的路由组件 默认是被**卸载**了的 需要的时候再进行**挂载**

#### `to`的两种写法

  ```html
  <!--第一种：字符串写法-->
<router-link active-class="active" to="/home">主页</router-link>
<!--第二种：to的对象写法-->
<router-link active-class="active" :to="{path:'/home'}">主页</router-link>
  ```

#### 路由的工作模式

##### 1. `history`模式

- **优点：** `URL`更加美观 不带有`#` 更接近传统网站的`URL`
- **缺点：** 后期项目上线 需要服务端配合处理路径问题 否则刷新会有`404`错误

```typescript
import {createWebHistory} from "vue-router";

const router = createRouter({
    history: createWebHistory(),  // 使用history模式
    // ...
})
```

##### 2. `hash`模式

- **优点：* 兼容性好 不需要服务端配合处理路径问题
- **缺点：** `URL`带有`#`不太美观 在`SEO`优化方面相对较差

```typescript
import {createWebHashHistory} from "vue-router";

const router = createRouter({
    history: createWebHashHistory(),  // 使用hash模式
    // ...
})
```

#### `to`的两种写法

##### 1.使用字符串写法

```html

<RouterLink to="/news" active-class="active">新闻</RouterLink>
```

##### 2.使用对象写法

```html

<RouterLink :to="{path: '/about'}" active-class="active">关于</RouterLink>
```

#### 路由传递参数

##### 1.query参数

- 什么是query参数
  ```
  http://localhost:5173/#/news/detail?id=000001&title=好消息&content=快过年了
  ```
- 传递参数

  ```html
  <li v-for="news in newsList" :key="news.id">
    <!--      query传递参数-->
    <RouterLink :to="{
          path: '/news/detail', // 跳转路径 也可以使用name
          query: {  // 配置query参数
            id: news.id,    // 主要配置项
            title: news.title,
            content: news.content,
          },
        }"> {{ news.title }}
    </RouterLink>
  </li>
  ```

- 接收参数

  ```typescript
  import { useRoute } from "vue-router";
  import { toRefs } from "vue";
  
  let route = useRoute()
  console.log(route)
  // 从一个响应式对象中直接解构出其属性 会使其丢失响应式 所以需要使用toRefs
  let { query } = toRefs(route)
  ```

#### params参数

- 传递params参数需要在路由配置文件`@/router/index.ts`中设置占位符
- `to`中不可以使用`path` 应该使用`name`

##### 声明占位符和设置name属性

  ```typescript
  const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        // ...
        {
            name: 'xinwen',
            path: '/news',
            component: News,
            children: [
                {
                    name: 'xiang',
                    // 使用params传参 需要进行占位 '?'表示此参数不是必要的
                    path: 'detail/:id/:title/:content?',
                    component: Detail,
                }
            ]
        },
        // ...
    ]
})
  ```

- 传递参数

  ```html
  <li v-for="news in newsList" :key="news.id">
    <!--      params传递参数-->
    <RouterLink :to="{
          name: 'xiang',  // 注意必须使用name 不能用path
          params: {
            id: news.id,
            title: news.title,
            content: news.content,
          },
        }"> {{ news.title }}
    </RouterLink>
  </li>
  ```

- 接收参数

  ```typescript
  import { useRoute } from "vue-router";
  import {toRefs} from "vue";
  let route = useRoute()
  let {params} = toRefs(route)
  ```

#### 路由的`props`配置

- **作用：** 让路由组件更方便的收到参数(可以将路由参数作为props传给组件)

  ```typescript
  // 路由规则启用 写法1 将路由组件收到所有params参数作为props传给子组件
  props: true,
  // 写法2 函数写法 可以自行设置什么作为props给路由组件
  props(route) {
    return route.query
  }
  ,
  // 写法3 对象写法 可以自行设置什么作为props给路由组件
  props: {
      a: 100,
  }
  ```

#### 编程式路由导航

- **作用：** 脱离`<RouterLink>`实现路由跳转

```typescript
interface NewInter {  // 对下面的news进行类型限时
    id?: string,
    title?: string,
    content?: string,
}

const router = useRouter()
let showNewsDetail = (news: NewInter) => {
    // 使用push方法 并携带参数
    router.push({
        name: 'xiang',
        query: {
            id: news.id,
            title: news.title,
            content: news.content
        }
    })
}
```

### 集中式状态(数据)管理 `pinia`

#### 安装`pinia`

- `npm install pinia`

#### 搭建`pinia`开发环境

##### 在`main.ts`中引入

```typescript
import {createPinia} from "pinia";  // 1. 引入pinia

const app = createApp(App)  // 创建一个应用

const pinia = createPinia() // 2. 创建pinia
app.use(pinia)  // 3. 安装pinia

app.mount('#app')
```

#### 创建数据存储 `@/store/`

```typescript
import {defineStore} from "pinia";
// defineStore中的参数<'文件名'>, {<配置对象>}
const useCountStore = defineStore('count', {
    state: () => {  // state 需要是一个函数 有返回值
        return {
            sum: 6,
        }
    }
})

export default useCountStore;
```

#### 在组件中读取数据

```typescript
import useCountStore from '@/store/count'

const countStore = useCountStore()

// 取值
countStore.sum
```

#### 修改数据的三种方式

1. 直接修改

```typescript
let add = () => countStore.sum += n.value
```

2. 碎片修改

```typescript
countState.$patch({
    sum: 888,
    address: 'JiangSu Changzhou'
})
```

3. 使用Action

- 先在`xxx.ts`中定义`action`

  ```typescript
  actions: {
      increment(val){
          console.log('increment被调用了', val, this);    // 这里的this就是state
          if (this.sum < 10)  // 假设这里是复杂的业务逻辑
              this.sum += val;
      }
  }
  ```

- 在组件中调用

  ```typescript
  countStore.increment(n.value)
  ```

#### `storeToRefs`的使用

- 借助`storeToRefs`将`store`中的数据转为`ref`对象 方便在模版中使用
- **注意：** `pinia`提供的`storeToRefs`只会将数据做转换 而`vue`的`toRefs`会转换`store`中所有的数据
  ```typescript
  import useCountStore from '@/store/count'
  // 引入
  import {storeToRefs} from "pinia";
  
  const countStore = useCountStore()
  // storeToRefs只会包裹数据的响应式
  let {sum, school, address} = storeToRefs(countStore)
  ```

#### `getters`的使用

- `Getter` 完全等同于 `Store` 状态的 计算值。

```typescript
getters: {
    bigSum: state => state.sum * 10, // 简写
        // bigSchool: state => state.school.toUpperCase(),
        bigSchool()
:
    string
    {
        return this.school.toUpperCase()
    }
}
```

#### `$subscribe` 的使用

- `subscriptions` 在 `patch` 后只触发一次

  ```typescript
  loveTalk.$subscribe((mutation, state) => {
    console.log('loveTalk里面保存的数据发生了变化', mutation, state);
    localStorage.setItem('talkList', JSON.stringify(state.talkList));
  })
  ```

```typescript
state()
{
    return {
        // talkList: [
        // {id: '000001', title: '今天你有点怪，哪里怪，怪好看的。'},
        // {id: '000002', title: '草莓、蓝莓、蔓越莓，今天想我了没？'},
        // {id: '000003', title: '心里给你留了一块地，我的死心塌地。'},
        // ]
        // 直接从浏览器本地存储中取出
        talkList: JSON.parse(localStorage.getItem('talkList') as string) || []
    }
}
,
```

#### `store`组合式的写法

```typescript
// 使用组合式
const useLoveTalk = defineStore('loveTalk', () => {
    // 这里的talkList就是state
    const talkList = reactive(JSON.parse(localStorage.getItem('talkList') as string) || []);

    // getATalk相当于actions
    async function getATalk() {
        // 发送请求 连续解构赋值和重命名
        let {data: {content: title}} = await axios.get('https://api.uomg.com/api/rand.qinghua?format=json')
        // 包装对象
        let obj = {
            id: nanoid(),
            title
        };
        // talkList.unshift(obj)
        talkList.unshift(obj)
    }

    return {talkList, getATalk}
})
```

### 组件间通信

#### **`Vue3`组件通信和`Vue2`的区别：**

- 移出事件总线 使用`mitt`代替
- `vuex`换成了`pini
- 把`.sync`优化到了`v-model`里面了
- 把`$listeners`所有的东西 合并到`$attrs`中了
- `$children`被砍掉了

#### props

- 概述：`props`是使用频率最高的一种通信方式，常用与 ：**父 ↔ 子**。
- 若 **父传子**：属性值是**非函数**。
- 若 **子传父**：属性值是**函数**。

父组件：

```vue

<template>
  <div class="father">
    <h3>父组件，</h3>
    <h4>我的车：{{ car }}</h4>
    <h4>儿子给的玩具：{{ toy }}</h4>
    <Child :car="car" :getToy="getToy"/>
  </div>
</template>

<script setup lang="ts" name="Father">
  import Child from './Child.vue'
  import {ref} from "vue";
  // 数据
  const car = ref('奔驰')
  const toy = ref()

  // 方法
  function getToy(value: string) {
    toy.value = value
  }
</script>
```

子组件

```vue

<template>
  <div class="child">
    <h3>子组件</h3>
    <h4>我的玩具：{{ toy }}</h4>
    <h4>父给我的车：{{ car }}</h4>
    <button @click="getToy(toy)">玩具给父亲</button>
  </div>
</template>

<script setup lang="ts" name="Child">
  import {ref} from "vue";

  const toy = ref('奥特曼')

  defineProps(['car', 'getToy'])
</script>
```

#### 自定义事件

1. 概述：自定义事件常用于：**子 => 父。**
2. 注意区分好：原生事件、自定义事件。

- 原生事件：
    - 事件名是特定的（`click`、`mosueenter`等等）
    - 事件对象`$event`: 是包含事件相关信息的对象（`pageX`、`pageY`、`target`、`keyCode`）
- 自定义事件：
    - 事件名是任意名称
    - <strong style="color:red">事件对象`$event`: 是调用`emit`时所提供的数据，可以是任意类型！！！</strong >

```html
   <!--在父组件中，给子组件绑定自定义事件：-->
<Child @send-toy="toy = $event"/>

<!--注意区分原生事件与自定义事件中的$event-->
<button @click="toy = $event">测试</button>
   ```

   ```js
   //子组件中，触发事件：
this.$emit('send-toy', 具体数据)
   ```

#### mitt

- **概述：** 与消息订阅与发布（`pubsub`）功能类似 可以实现任意组件间通信
- 安装`mitt`

  ```shell
  npm i mitt
  ```

- 新建文件：`src\utils\emitter.ts`

  ```javascript
  // 引入mitt 
  import mitt from "mitt";
  
  // 创建emitter
  const emitter = mitt()
  
  /*
    // 绑定事件
    emitter.on('abc',(value)=>{
      console.log('abc事件被触发',value)
    })
    emitter.on('xyz',(value)=>{
      console.log('xyz事件被触发',value)
    })
  
    setInterval(() => {
      // 触发事件
      emitter.emit('abc',666)
      emitter.emit('xyz',777)
    }, 1000);
  
    setTimeout(() => {
      // 清理事件
      emitter.all.clear()
    }, 3000); 
  */
  
  // 创建并暴露mitt
  export default emitter
  ```

- 接收数据的组件中：绑定事件 同时在销毁前解绑事件

  ```typescript
  import emitter from "@/utils/emitter";
  import { onUnmounted } from "vue";
  
  // 绑定事件
  emitter.on('send-toy',(value)=>{
    console.log('send-toy事件被触发',value)
  })
  
  onUnmounted(()=>{
    // 解绑事件
    emitter.off('send-toy')
  })
  ``` 

#### v-model

- **概述：** 实现 **父↔子** 之间相互通信。

- 前序知识 —— `v-model`的本质

   ```vue
   <!-- 使用v-model指令 -->
   <input type="text" v-model="userName">
   
   <!-- v-model的本质是下面这行代码 -->
   <input 
     type="text" 
     :value="userName" 
     @input="userName =(<HTMLInputElement>$event.target).value"
   >
   ```

3. 组件标签上的`v-model`的本质：`:moldeValue` ＋ `update:modelValue`事件。

   ```vue
   <!-- 组件标签上使用v-model指令 -->
   <AtguiguInput v-model="userName"/>
   
   <!-- 组件标签上v-model的本质 -->
   <AtguiguInput :modelValue="userName" @update:model-value="userName = $event"/>
   ```

   `AtguiguInput`组件中：

   ```vue
   <template>
     <div class="box">
       <!--将接收的value值赋给input元素的value属性，目的是：为了呈现数据 -->
   		<!--给input元素绑定原生input事件，触发input事件时，进而触发update:model-value事件-->
       <input 
          type="text" 
          :value="modelValue" 
          @input="emit('update:model-value',$event.target.value)"
       >
     </div>
   </template>
   
   <script setup lang="ts" name="AtguiguInput">
     // 接收props
     defineProps(['modelValue'])
     // 声明事件
     const emit = defineEmits(['update:model-value'])
   </script>
   ```

4. 也可以更换`value`，例如改成`abc`

   ```vue
   <!-- 也可以更换value，例如改成abc-->
   <AtguiguInput v-model:abc="userName"/>
   
   <!-- 上面代码的本质如下 -->
   <AtguiguInput :abc="userName" @update:abc="userName = $event"/>
   ```

   `AtguiguInput`组件中：

   ```vue
   <template>
     <div class="box">
       <input 
          type="text" 
          :value="abc" 
          @input="emit('update:abc',$event.target.value)"
       >
     </div>
   </template>
   
   <script setup lang="ts" name="AtguiguInput">
     // 接收props
     defineProps(['abc'])
     // 声明事件
     const emit = defineEmits(['update:abc'])
   </script>
   ```

5. 如果`value`可以更换，那么就可以在组件标签上多次使用`v-model`

   ```vue
   <AtguiguInput v-model:abc="userName" v-model:xyz="password"/>
   ```

#### $attrs

- **概述：** `$attrs`用于实现**当前组件的父组件**，向**当前组件的子组件**通信（**祖→孙**）。
- **具体说明：** `$attrs`是一个对象，包含所有父组件传入的标签属性。

  > 注意：`$attrs`会自动排除`props`中声明的属性(可以认为声明过的 `props` 被子组件自己“消费”了)

- 父组件：

  ```vue
  <template>
    <div class="father">
      <h3>父组件</h3>
          <Child :a="a" :b="b" :c="c" :d="d" v-bind="{x:100,y:200}" :updateA="updateA"/>
    </div>
  </template>
  
  <script setup lang="ts" name="Father">
      import Child from './Child.vue'
      import { ref } from "vue";
      let a = ref(1)
      let b = ref(2)
      let c = ref(3)
      let d = ref(4)
  
      function updateA(value){
          a.value = value
      }
  </script>
  ```

- 子组件：

  ```vue
  <template>
      <div class="child">
          <h3>子组件</h3>
          <GrandChild v-bind="$attrs"/>
      </div>
  </template>
  
  <script setup lang="ts" name="Child">
      import GrandChild from './GrandChild.vue'
  </script>
  ```

- 孙组件：

  ```vue
  <template>
      <div class="grand-child">
          <h3>孙组件</h3>
          <h4>a：{{ a }}</h4>
          <h4>b：{{ b }}</h4>
          <h4>c：{{ c }}</h4>
          <h4>d：{{ d }}</h4>
          <h4>x：{{ x }}</h4>
          <h4>y：{{ y }}</h4>
          <button @click="updateA(666)">点我更新A</button>
      </div>
  </template>
  
  <script setup lang="ts" name="GrandChild">
      defineProps(['a','b','c','d','x','y','updateA'])
  </script>
  ```

#### $refs 和 $parent

- 概述：

    * `$refs`用于 ：**父→子。**
    * `$parent`用于：**子→父。**

2. 原理如下：

   | 属性      | 说明                                                     |
         | --------- | -------------------------------------------------------- |
   | `$refs`   | 值为对象，包含所有被`ref`属性标识的`DOM`元素或组件实例。 |
   | `$parent` | 值为对象，当前组件的父组件实例对象。                     |

### 其他API

#### `shallowRef` & `shallowReactive`

##### `shallowRef`

- **作用：** 创建一个响应式数据 但只对顶层属性进行响应式处理
- **用法：**
  ```typescript
  let myVar = shallowRef(initialValue);
  ```

- **特点：** 只跟踪引用值的变化 不关心值内部的属性变化

##### `shallowReactive`

- **作用：** 创建一个浅层响应式对象 只会使对象的最顶层属性变成响应式的 对象内部的嵌套属性则不会变成响应式的

- **用法：**

   ```js
   const myObj = shallowReactive({ ... });
   ```
- **特点：** 对象的顶层属性是响应式的 但嵌套对象的属性不是 

##### 总结

> 通过使用 [`shallowRef()`](https://cn.vuejs.org/api/reactivity-advanced.html#shallowref)
> 和 [`shallowReactive()`](https://cn.vuejs.org/api/reactivity-advanced.html#shallowreactive) 来绕开深度响应。浅层式 `API`
> 创建的状态只在其顶层是响应式的，对所有深层的对象不会做任何处理，避免了对每一个内部属性做响应式所带来的性能成本，这使得属性的访问变得更快，可提升性能。

#### `readonly` & `shallowReadonly`

##### **`readonly`**

- **作用：** 用于创建一个对象的深只读副本

- **用法：**

   ```js
   const original = reactive({ ... });
   const readOnlyCopy = readonly(original);
   ```

- **特点：**

  * 对象的所有嵌套属性都将变为只读
  * 任何尝试修改这个对象的操作都会被阻止（在开发模式下 还会在控制台中发出警告）

- **应用场景：**
  * 创建不可变的状态快照。
  * 保护全局状态或配置不被修改

##### **`shallowReadonly`**
- **作用：** 与 `readonly` 类似，但只作用于对象的顶层属性

- **用法：**

   ```js
   const original = reactive({ ... });
   const shallowReadOnlyCopy = shallowReadonly(original);
   ```

- **特点：**

  * 只将对象的顶层属性设置为只读，对象内部的嵌套属性仍然是可变的。

  * 适用于只需保护对象顶层属性的场景。

#### `toRaw` & `markRaw`

#### `customRef`

## 8.2. 【Suspense】

-  等待异步组件时渲染一些额外内容，让应用有更好的用户体验
-  使用步骤：
  -  异步引入组件
  -  使用`Suspense`包裹组件，并配置好`default` 与 `fallback`

```tsx
import { defineAsyncComponent,Suspense } from "vue";
const Child = defineAsyncComponent(()=>import('./Child.vue'))
```

```vue
<template>
    <div class="app">
        <h3>我是App组件</h3>
        <Suspense>
          <template v-slot:default>
            <Child/>
          </template>
          <template v-slot:fallback>
            <h3>加载中.......</h3>
          </template>
        </Suspense>
    </div>
</template>
```

[Note]: Kanned1079

