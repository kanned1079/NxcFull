// 項目的根組件
import "./App.css"

import {useEffect, useState} from "react";
// useEffect(() => {}, []);
// 参数1 是一个函数 可以把它叫做副作用函数 在函数内部可以放置要执行的操作】
// 参数2 是一个数组(可选参数) 在数组中放置依赖项目 不同依赖项会影响第一个函数的执行 **当是一个空数组的时候 副作用函数只会在组件渲染完毕后执行一次**

// 1.没有依赖项 初始+组件更新时
// 2.空数组依赖 初始执行一次
// 3.传入特定依赖项 初始+依赖项变化时

function App() {
    let [show, setShow] = useState(true);
    return (
        <div className="App">
            <h2>This is App</h2>
            <button
                className="son-send-btn"
                onClick={() => setShow(false)}
            >卸载Son组件</button>
            {show && <Son />}
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
