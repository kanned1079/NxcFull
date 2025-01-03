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

