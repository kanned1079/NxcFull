## React + TypeScripts + Vite + SWC 

### `useState` 

#### 自动推倒

- 通常React会根据传入`useState的默认值`来自动推导类型,不需要显式标注类型

```jsx
// 一般类型
const [value, setValue] = useState(false)
const changeValue = () => setValue(!value);
// 复杂类型
const [numList, setNumList] = useState([1, 2, 3]);
const changeList = () => setNumList([4, 5, 6])
```

#### 传递范型参数

- useState本身是一个`泛型函数` 可以传入具体的自定义类型

```jsx
// 定义类型
type MyUser = {
    name: string
    age: number
}
```

```jsx
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

```jsx
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