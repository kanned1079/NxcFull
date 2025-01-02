// 項目的根組件
import "./App.css"

import {useSelector, useDispatch} from "react-redux"
import {increment, decrement, add2Nums} from "./store/modules/counterStore.js";
import {fetchChannels} from "./store/modules/channelStore.js";
import {useEffect} from "react";

function App() {
    const {count} = useSelector(state => state.counter)
    const {channelList} = useSelector(state => state.channel)
    const dispatch = useDispatch()
    useEffect(() => {
        dispatch(fetchChannels())
    }, [dispatch]);
    return (
        <div className="App">
            <div>
                <p>count: {count}</p>
                <button onClick={() => dispatch(increment())}>increment</button>
                <button onClick={() => dispatch(decrement())}>decrement</button>
                <button onClick={() => dispatch(add2Nums(10))}>decrement to 10</button>
                <button onClick={() => dispatch(add2Nums(20))}>decrement to 20</button>
            </div>
            <div>
                <ul>
                    {channelList.map((item, index) => (<li key={item.id}>{item.name}</li>))}
                </ul>
            </div>
        </div>
    );
}


export default App;
