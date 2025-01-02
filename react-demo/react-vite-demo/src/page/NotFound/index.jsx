import React, {useEffect} from "react"
import ReactDOM from 'react-dom/client';
import "./index.less"
import {Link, useNavigate} from "react-router-dom";

function NotFound() {
    const navigate = useNavigate();
    useEffect(() => {
        console.log('aaa')
    }, []);
    return (
        <div className="root">
            <div className="login-card">
                <p className="hex">404 Not Found</p>
            </div>
        </div>
    )
}

export default NotFound
