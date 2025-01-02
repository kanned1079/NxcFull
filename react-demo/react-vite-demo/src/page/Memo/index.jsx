import {useMemo, useState, memo} from "react";

import styles from "./index.module.less"

const Memo = () => {
    let [count1, setCount1] = useState(0);
    let [count2, setCount2] = useState(0);
    let getFib = (n) => {
        console.log('执行计算函数')
        return  n < 3?1:getFib(n - 2) + getFib(n - 1)
    }
    let result1 = useMemo(() => {
        // 返回计算后的结果
        return getFib(count1)
    }, [count1])
    return (
        <div className={styles.root}>
            <h3>fib: {result1}</h3>
            <button onClick={() => {setCount1(count1 + 1)}}>count1: {count1}</button>
            <button onClick={() => {setCount2(count2 + 1)}}>count2: {count2}</button>
            <Son />
            <MemoSon />
        </div>
    )
}

const Son = () => {
    console.log('Son渲染了')
    return (
        <div>MemoSon</div>
    )
}

const MemoSon = memo(function Son () {
    console.log('MemoSon渲染了')
    return (
        <div>MemoSon</div>
    )
})

export default Memo