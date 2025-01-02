import React, {useEffect} from "react"
import ReactDOM from 'react-dom/client';
import "./index.less"

function Login() {
    useEffect(() => {
        console.log('aaa')
    }, []);
    return (
        <div className="root">
            <div className="login-card">Login</div>
        </div>
    )
}

export default Login