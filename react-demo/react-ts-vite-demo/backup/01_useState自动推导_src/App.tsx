import {useState} from "react";

function App() {
    const [value, setValue] = useState(false)
    const changeValue = () => setValue(!value);

    const [numList, setNumList] = useState([1, 2, 3]);
    const changeList = () => setNumList([4, 5, 6])

    return (
        <>
            <h2>This is app.</h2>
            <p>{value ? 'TES' : 'NO'}</p>
            <button onClick={() => changeValue()}>toggle</button>
            {numList.map((num: number, index: number) => (
                <p key={index}>{num}</p>
            ))}
            <button onClick={() => changeList()}>toggle</button>

        </>
    )
}

export default App
