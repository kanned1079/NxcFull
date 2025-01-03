import {memo, useCallback, useState} from "react"

const MyInput = memo(function MyInput (props) {
    console.log('子组件渲染了')
    return <input type="text" onChange={(e) => props.onChange(e.target.value)} />
})

const App = () => {
    const [count, setCount] = useState(0)
    const handleChange = useCallback((value) => console.log(value), [])
    return (
        <div>
            <MyInput onChange={handleChange} />
            <button onClick={() => setCount(count + 1)}>{count}</button>
        </div>
    )
}

export default App