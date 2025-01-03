import React from "react"
import ReactDOM from 'react-dom/client';
import "./index.less"
import {Link, useParams, useSearchParams} from "react-router-dom";

const Article = () => {
    // const [params] = useSearchParams();
    // let id = params.get('id')
    // let name = params.get('name')
    const params = useParams()
    let id = params.id
    let name = params.name
    return (
        <div className="root">
            <div className="login-card">
                <p>article</p>
                <p>{id}</p>
                <p>{name}</p>
            </div>
        </div>
    )
}

export default Article