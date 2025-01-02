// 項目的根組件
import {useState} from "react";

function App() {
    const [count, setCount] = useState(0)
    // 第一個是值 第二個是修改值的函數
    const handleAdd = () => {
        // setCount 用新值設置count 並使用新的count渲染dom
        setCount(count + 1)
    }

    return (
        <div className="App">
            <button onClick={handleAdd}>+1</button>
            <p>{count}</p>
        </div>
    );
}

export default App;
