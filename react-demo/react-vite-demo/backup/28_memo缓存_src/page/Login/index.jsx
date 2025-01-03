import React, {useEffect} from "react"
import ReactDOM from 'react-dom/client';
import "./index.less"
import {Link, useNavigate} from "react-router-dom";

function Login() {
    const navigate = useNavigate();
    useEffect(() => {
        console.log('aaa')
    }, []);
    return (
        <div className="root">
            <div className="login-card">
                Login
                {/*声明式方法*/}
                <Link to={'/article'}>跳转到到Article</Link>
                {/*命令式的方法*/}
                <button onClick={() => {
                    navigate('/article?id=101&name=kanna')
                }}>searchParams传参
                </button>
                <button onClick={() => {
                    navigate('/article/101/kanna')
                }}>params传参
                </button>
            </div>
        </div>
    )
}

export default Login