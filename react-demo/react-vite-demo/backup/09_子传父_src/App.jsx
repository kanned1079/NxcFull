// 項目的根組件
import "./App.css"
import {useState} from "react"

function App() {
    let [str, setStr] = useState("");
    const getMsg = (temp) => {
        setStr(temp)
    }
    return (
        <div className="App">
            <h2>This is App</h2>
            <p>msg: {str}</p>
            <Son onGetMsg={getMsg}></Son>
        </div>
    );
}

// 子 -> 父
// 核心：在自组件中调用父组件中的函数并传递参数

function Son(props) {
    const message = 'this is message from son'

    return (
        <div className="Son" style={{backgroundColor: '#e3e5e7'}}>
            <h2>This is son</h2>
            <button onClick={() => {
                props.onGetMsg(message)
            }} className="son-send-btn">send</button>
        </div>
    )
}

export default App;
