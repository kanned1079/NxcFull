// 項目的根組件
import "./App.css"

import React, {useEffect, useState} from "react";

// ReactHooks使用规则
// 1.只能在组件中或者其他自定义Hook函数中调用
// const [val, setVal] = useState("");
// 2.只能在组件的顶层调用，不能嵌套在if、for、其他函数中
// function TestComponent() {
//     if (true) {
//         let [value, setValue] = useState(0);
//     }
//     for (let i = 0; i < 10; i++) {
//         let [value, setValue] = useState(0);
//     }
// }



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
