// import Login from "../page/Login"
// import Article from "../page/Article"

import Layout from "../page/Layout"
import About from "../page/About"
import Board from "../page/Board"

import {createBrowserRouter} from "react-router-dom";

const router = createBrowserRouter([
    {
        path: '/',
        element: <Layout/>,
        children: [
            {
                path: '/about',
                element: <About/>
            },
            {
                path: '/board',
                element: <Board/>
            },
        ]
    },
])

export default router