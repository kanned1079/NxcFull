import React, {useEffect} from "react"
import ReactDOM from 'react-dom/client';
import {Link, useNavigate} from "react-router-dom";

function NotFound() {
    const navigate = useNavigate();
    useEffect(() => {
        console.log('aaa')
    }, []);
    return (
        <div className={styles.root}>
            <div className={styles.loginCard}>
                <p className={styles.hex}>404 Not Found</p>
            </div>
        </div>
    )
}

export default NotFound
