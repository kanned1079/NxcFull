// 項目的根組件
import "./App.css"

import {useEffect, useState} from "react";
// useEffect(() => {}, []);
// 参数1 是一个函数 可以把它叫做副作用函数 在函数内部可以放置要执行的操作】
// 参数2 是一个数组(可选参数) 在数组中放置依赖项目 不同依赖项会影响第一个函数的执行 **当是一个空数组的时候 副作用函数只会在组件渲染完毕后执行一次**

const Url = 'http://geek.itheima.net/v1_0/channels'

function App() {
    const [list, setList] = useState([])
    useEffect(async () => {
        // 这里是额外的操作
        let getList = async () => {
            let response = await fetch(Url)
            const {data} = await response.json()
            setList(data.channels)
        }
        await getList()
    }, []); // **当是一个空数组的时候 副作用函数只会在组件渲染完毕后执行一次**
    return (
        <div className="App">
            <h2>This is App</h2>
            <ul>
                {list.map(item => <li key={item.id}>{item.name}</li>)}
            </ul>
        </div>
    );
}

export default App;
