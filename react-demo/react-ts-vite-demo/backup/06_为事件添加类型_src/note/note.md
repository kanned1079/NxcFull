## React + TypeScripts + Vite + SWC 

### `useState` 

#### 自动推倒

- 通常React会根据传入`useState的默认值`来自动推导类型,不需要显式标注类型

```tsx
// 一般类型
const [value, setValue] = useState(false)
const changeValue = () => setValue(!value);
// 复杂类型
const [numList, setNumList] = useState([1, 2, 3]);
const changeList = () => setNumList([4, 5, 6])
```

#### 传递范型参数

- useState本身是一个`泛型函数` 可以传入具体的自定义类型

```tsx
// 定义类型
type MyUser = {
    name: string
    age: number
}
```

```tsx
// 限制初始类型
    const [user1, setUser1] = useState<MyUser>({	// 对象型
        name: 'kinggyo',
        age: 16,
    });
    const [user2, setUser2] = useState<MyUser>(() => {	// 函数返回型
        return {
            name: 'kanna',
            age: 16,
        }
    });
```

#### 修改

```tsx
const changeUser = () => {
        // setUser1((): MyUser => ({	// 函数返回型
        //     name: 'kinggyo2',
        //     age: 18,
        // }))
        setUser1({		// 对象型
            name: 'kinggyo2',
            age: 18,
        })
        // setTimeout(undefined)    // 如果前面没有初始值可以用
    }
```

#### `useState`初始值为`null`

- 当我们不知道状态的初始值是什么 将usestate的`初始值为nul`是一个常见的做法 可以通过`具体类型联合null`来做显
  式注解

```tsx
const [user1, setUser1] = useState<MyUser | null>(null);
const changeUser = () => {
        // setUser1(null);     // 可以设置为null
        setUser1({
            name: 'kinggyo2',
            age: 18,
        })
    }
```

### `Perops` && `TypeScripts`

```tsx
interface MyProps {
    addr: string
    title?: string
}

const Son = (props: MyProps) => {
    return (
        <>
            <h2>{props.title}</h2>
            <h3> Addr: {props.addr} </h3>
        </>
    )
}
```

#### 特殊的Props 

- children是一个比较特殊的prop 支持多种不同类型数据的传入 需要通过一个`内置的ReactNode类型`来做注解

```tsx
interface MyProps {
    addr: string
    title?: string
    children?: ReactNode
}

const Son = (props: MyProps) => {
    return (
        <>
            <h2>{props.title}</h2>
            <h3> Addr: {props.addr} </h3>
            {props.children}
        </>
    )
}
```

#### 为事件prop添加类型

```tsx
interface MyProps {
    addr: string
    children?: ReactNode
    onSetUser?: (name: string, age: number) => void
}
```

```tsx
{/* 调用 */}
<button onClick={() => {props.onSetUser?.('kanna', 22)}}>调用方法</button>
```

