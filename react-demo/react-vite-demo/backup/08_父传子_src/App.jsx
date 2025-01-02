// 項目的根組件
import "./App.css"
import {useState} from "react"

function App() {
    const name = 'kinggyo'
    return (
        <div className="App">
            <h2>This is App</h2>
            {/*props能传递几乎所有数据*/}
            <Son
                name={name}
                age={18}
                isTrue={true}
                list={['vue', 'react', 'angular']}
                obj={{
                    name: 'kanna',
                    age: 21,
                }}
                childDom={<div style={{fontSize: '1rem', color: 'red'}}>传递的JSX</div>}
                handle={() => console.log('传递的函数')}
            ></Son>

            {/*一个特殊的props*/}
            <Son2>
                <span>this is span</span>
                {/*这样传递后在子组件中的props中会自动多一个children属性*/}
            </Son2>

        </div>
    );
}

// 父 -> 子
// 1.子组件传递数据 自组件身上绑定属性
// 2. 子组件接收数据 props的参数

function Son(props) {
    console.log(props)  // 这里的props是只读的
    props.handle()  // 函数
    // props.name = 'aaa'   // 不能修改 否则会报错
    return (
        <div className="Son" style={{backgroundColor: '#e3e5e7'}}>
            <h2>This is son</h2>
            <h3>Name {props.name}</h3>
            {props.list.map(item => <li style={{color: 'blue'}} key={item}>{item}</li>)}
            {props.childDom}
        </div>
    )
}

// 一个特殊的props
// 当我们把内容嵌套在自组件标签中时 父组件会自动在名为children的props属性中接收该内容
function Son2(props) {
    console.log(props)
    return (
        <div className="Son2">
            <h2>This is son2</h2>
            {props.children}
        </div>
    )
}

export default App;
