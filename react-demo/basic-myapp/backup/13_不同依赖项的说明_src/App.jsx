// 項目的根組件
import "./App.css"

import {useEffect, useState} from "react";
// useEffect(() => {}, []);
// 参数1 是一个函数 可以把它叫做副作用函数 在函数内部可以放置要执行的操作】
// 参数2 是一个数组(可选参数) 在数组中放置依赖项目 不同依赖项会影响第一个函数的执行 **当是一个空数组的时候 副作用函数只会在组件渲染完毕后执行一次**

// 1.没有依赖项 初始+组件更新时
// 2.空数组依赖 初始执行一次
// 3.传入特定依赖项 初始+依赖项变化时

const Url = 'http://geek.itheima.net/v1_0/channels'

function App() {
    let [count, setCount] = useState(0);
    let [list, setList] = useState([])
    const getList = async () => {
        let response = await fetch(Url)
        const {data} = await response.json()
        setList(data.channels)
    }
    // 初始+组件更新时
    // useEffect(() => {   // 这里是额外的操作
    //     console.log('副作用函数执行了')
    // }, []);
    // 初始执行一次
    // useEffect(() => {   // 这里是额外的操作
    //     console.log('副作用函数执行了')
    // }, []);
    // 初始+依赖项变化时
    useEffect(() => {   // 这里是额外的操作
        console.log('副作用函数执行了', count)
    }, [count]);    // 和第一个相比 这个只有当count变化的时候才会执行
    return (
        <div className="App">
            <h2>This is App</h2>
            <button
                className="son-send-btn"
                onClick={() => setCount(count + 1)}
            >更新组件 {count}</button>
            <ul>
                {list.map(item => <li key={item.id}>{item.name}</li>)}
            </ul>
        </div>
    );
}

export default App;
