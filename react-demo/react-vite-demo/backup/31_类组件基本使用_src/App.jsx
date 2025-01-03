import {Component, forwardRef, memo, useImperativeHandle, useRef} from "react"

const MyInput = memo(function MyInput(props) {
    console.log('子组件渲染了')
    return <input type="text" onChange={(e) => props.onChange(e.target.value)}/>
})

const MyInput2 = forwardRef((props, ref) => {
    const inputRef = useRef(null)
    const handleFocus = () => { // 注意focus是在子组件中的
        inputRef.current.focus()
    }
    // 暴露函数给父组件调用
    useImperativeHandle(ref, () => {
        return {
            handleFocus,
        }
    })
    return <input type="text" ref={inputRef}/>
})

class Counter extends Component {
    // 状态
    state = {
        count: 0,
    }
    // 事件回调修改状态
    setCount = () => {
        this.setState({
            count: this.state.count + 1,
        })
    }
    // 模版
    render() {
        return (
            <button onClick={this.setCount}>{this.state.count}</button>
        )
    }
}

const App = () => {
    const inputRef = useRef(null)
    const showRef = () => {
        console.log(inputRef.current)
        inputRef.current.handleFocus()
    }
    return (
        <>
            <MyInput2 ref={inputRef}/>
            <button onClick={showRef}>focus</button>
            <Counter></Counter>
        </>
    )
}

export default App