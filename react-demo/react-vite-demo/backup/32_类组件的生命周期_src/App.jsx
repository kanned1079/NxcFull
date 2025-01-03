import {Component, forwardRef, memo, useImperativeHandle, useRef, useState} from "react"

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

    componentDidMount() {
        console.log('组件挂载')
         this.intervalId = setInterval(() => {
            console.log('定时器输出')
        }, 200)
    }
    componentWillUnmount() {
        console.log('组件卸载')
        clearInterval(this.intervalId)
    }

    // 模版
    render() {
        return (
            <div>
                <br/>
                <button onClick={this.setCount}>{this.state.count}</button>
            </div>
        )
    }
}

const App = () => {
    const inputRef = useRef(null)
    const showRef = () => {
        console.log(inputRef.current)
        inputRef.current.handleFocus()
    }
    const [show, setShow] = useState(true)
    return (
        <>
            <MyInput2 ref={inputRef}/>
            <button onClick={showRef}>focus</button>
            { show && <Counter></Counter> }
            <button onClick={() => setShow(!show)}>卸载Counter</button>
        </>
    )
}

export default App