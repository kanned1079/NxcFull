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
    const myCount = 100
    const myLists = [1, 2, 3]
    const myLists2 = useMemo(() => {
        return [1, 2, 3]
    }, [])
    return (
        <div className={styles.root}>
            <h3>fib: {result1}</h3>
            <button onClick={() => {setCount1(count1 + 1)}}>count1: {count1}</button>
            <button onClick={() => {setCount2(count2 + 1)}}>count2: {count2}</button>
            <Son />
            <MemoSon count={myCount}></MemoSon>
            <MemoSonComplex list={myLists}></MemoSonComplex>
            <MemoSonComplex list={myLists2}></MemoSonComplex>
        </div>
    )
}

const Son = () => {
    console.log('Son渲染了')
    return (
        <div>MemoSon</div>
    )
}

// 演示memo的简单类型
const MemoSon = memo(function Son (props) {
    console.log('MemoSon渲染了')
    return (
        <div>
            MemoSon
            <p>son's count: {props.count}</p>
        </div>

    )
})

// 演示memo的复杂类型
const MemoSonComplex = memo(function Son (props) {
    console.log('MemoSonComplex2渲染了')
    return (
        <div>
            MemoSon
            <p>son's count: {props.list}</p>
        </div>

    )
})

export default Memo