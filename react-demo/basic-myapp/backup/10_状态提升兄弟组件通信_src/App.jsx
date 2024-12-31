// 項目的根組件
import "./App.css"
import {useState} from "react"

// 1.子->父 A->App
// 2.父传子 App->B

function App() {
    let [name, setName] = useState("");
    const getName = (temp) => setName(temp);
    return (
        <div className="App">
            <h2>This is App</h2>
            <SonA onSetName={setName}></SonA>
            <SonB name={name}></SonB>
        </div>
    );
}

// 子 -> 父
// 核心：在自组件中调用父组件中的函数并传递参数

function SonA(props) {
    const name = 'kinggyo'
    return (
        <div className="Son" style={{backgroundColor: '#e3e5e7'}}>
            <h2>This is sonA</h2>
            <button
                onClick={() => props.onSetName(name)}
                className="son-send-btn"
            >Send to SonB</button>
        </div>
    )
}

function SonB(props) {
    return (
        <div className="Son" style={{backgroundColor: '#e3e5e7'}}>
            <h2>This is sonB</h2>
            <p className="result">name: {props.name}</p>
        </div>
    )
}

export default App;
