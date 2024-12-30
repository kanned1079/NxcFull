// 項目的根組件
import {useState} from "react";

function App() {
    // 在React中 狀態是只讀的 應該替換它而不是修改它
    const [count, setCount] = useState(0)
    // 第一個是值 第二個是修改值的函數
    const handleAdd = () => {
        // setCount 用新值設置count 並使用新的count渲染dom
        // count++  // 這樣是不行的
        setCount(count + 1) // 要用這個
    }

    let [form, setForm] = useState({
        name: 'kinggyo',
    })
    let handleChangeName = () => {
        // form.name = 'kanna' // 雖然修改了但是dom是沒有更新的
        console.log(form.name)
        setForm({
            ...form,    // 先展其中的其他字段 如果有
            name: 'kanna',   // 再修改要更改的字段
        })

    }

    return (
        <div className="App">
            <h4>一般數據</h4>
            <button onClick={handleAdd}>+1</button>
            <p>{count}</p>
            <h4>對象數據</h4>
            <button onClick={handleChangeName}>改為kanna</button>
            <p>{form.name}</p>
        </div>
    );
}

export default App;
