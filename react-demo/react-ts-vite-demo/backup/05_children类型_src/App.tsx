import {useState, type ReactNode} from "react";

type MyUser = {
    name: string
    age: number
}

function App() {
    // 限制初始类型
    const [user1, setUser1] = useState<MyUser | null>(null);
    // const [user2, setUser2] = useState<MyUser>(() => {
    //     return {
    //         name: 'kanna',
    //         age: 16,
    //     }
    // });
    // 修改的两种形式
    const changeUser = () => {
        // setUser1(null);     // 可以设置为null
        setUser1({
            name: 'kinggyo2',
            age: 18,
        })
    }

    return (
        <>
            <h2>This is app.</h2>
            <div style={{backgroundColor: '#e3e5e7', padding: 20}}>
                <h3>{user1?.name}</h3>
                <h4>{user1?.age}</h4>
            </div>
            <button onClick={() => changeUser()}>修改</button>
            <Son addr={'上海'} title={'这是标题'}>
                <h4>这是插槽的内容</h4>
            </Son>

        </>
    )
}

interface MyProps {
    addr: string
    title?: string
    children?: ReactNode
}

const Son = (props: MyProps) => {
    return (
        <>
            <h2>{props.title}</h2>
            <h3> Addr: {props.addr} </h3>
            {props.children}
        </>
    )
}

export default App
