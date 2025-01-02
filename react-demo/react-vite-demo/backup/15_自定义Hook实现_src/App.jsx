// 項目的根組件
import "./App.css"

import {useEffect, useState} from "react";

// boolean切换的逻辑是和当前组件耦合在一起的 不方便使用
// 解决思路：自定义Hook
// 1.声明一个以use起头的函数
// 2.在函数题中封装可以复用的逻辑(只要是可以复用的逻辑)
// 3.把组件中需要用到的状态或回调进行return(以对象或数组)
// 4.在哪个组件中要用到这个逻辑 就执行这个函数并结构出来进行使用

let useToggle = () => {
    // 这里是可以复用的逻辑
    let [show, setShow] = useState(true);
    const toggle = () => setShow(!show)
    // 哪些状态和函数需要在其他组件中使用 使用return
    return {
        show, toggle
    }
}

function App() {
    // let [show, setShow] = useState(true);
    // const toggle = () => setShow(!show)
    const {show, toggle} = useToggle()

    return (
        <div className="App">
            <h2>This is App</h2>
            <button
                className="son-send-btn"
                onClick={toggle}
            >toggle</button>
            {show && <div className="SonA">aaa</div>}
        </div>
    );
}

function Son() {
    // 1.渲染完毕后开启一个定时器
    useEffect(() => {
        let intervalId = setInterval(() => {
            console.log('定时器执行中...')
        }, 1000)
        // 组件卸载的时候清楚定时器
        return () => {
            console.log('清除定时器')
            clearInterval(intervalId)
        }
    }, []);
    return (
        <div className="SonA">
            <h2>This is Son</h2>
        </div>
    )
}

export default App;
