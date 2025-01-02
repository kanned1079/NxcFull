// 項目的根組件
import {useState} from "react";
import "./App.css"

function App() {

    return (
        <div className="App">
            <div style={{color: 'red', fontSize: '1.5rem'}}>行內樣式控制1</div>
            <div style={myStyle}>行內樣式控制2</div>
            <hr/>
            <div className="root">通過類名指定</div>
        </div>
    );
}

const myStyle = {
    color: 'blue',
    fontSize: '1rem',
}

export default App;
