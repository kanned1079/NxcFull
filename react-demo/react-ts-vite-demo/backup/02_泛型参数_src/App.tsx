import {useState} from "react";

type MyUser = {
    name: string
    age: number
}

function App() {
    // 限制初始类型
    const [user1, setUser1] = useState<MyUser>({
        name: 'kinggyo',
        age: 16,
    });
    // const [user2, setUser2] = useState<MyUser>(() => {
    //     return {
    //         name: 'kanna',
    //         age: 16,
    //     }
    // });
    // 修改的两种形式
    const changeUser = () => {
        // setUser1((): MyUser => ({
        //     name: 'kinggyo2',
        //     age: 18,
        // }))
        setUser1({
            name: 'kinggyo2',
            age: 18,
        })
        // setTimeout(undefined)    // 如果前面没有初始值可以用
    }

    return (
        <>
            <h2>This is app.</h2>
            <div style={{backgroundColor: '#e3e5e7', padding: 20}}>
                <h3>{user1?.name}</h3>
                <h4>{user1?.age}</h4>
            </div>
            <button onClick={() => changeUser()}>修改</button>


        </>
    )
}

export default App
