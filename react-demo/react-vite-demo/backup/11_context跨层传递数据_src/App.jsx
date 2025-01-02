// 項目的根組件
import "./App.css"
import {useState, createContext, useContext} from "react"

// 1.使用createContext方法创建一个上下文对象Ctx
const MsgContext = createContext()
// 2.在顶层组件App中通过Ctx.Provider组件提供数据
// 3.在底层组件B中通过useContext钩子函数获取消息数据

function App() {
    const msg = 'this is App\'s msg'
    return (
        <div className="App">
            <h2>This is App</h2>
            <MsgContext.Provider
                value={msg}
            >
                <SonA></SonA>
            </MsgContext.Provider>


        </div>
    );
}

// 跨层通信
// 使用Context机制跨层级组件通信

function SonA() {
    return (
        <div className="SonA">
            <h2>This is sonA</h2>
            <SonB></SonB>
        </div>
    )
}

function SonB(props) {
    const msg = useContext(MsgContext)
    return (
        <div className="SonB">
            <h2>This is sonB</h2>
            <p className="result">msg: {msg}</p>
        </div>
    )
}

export default App;
