// import Login from "../page/Login"
// import Article from "../page/Article"

import Layout from "../page/Layout"
import About from "../page/About"
import Board from "../page/Board"
import NotFound from "../page/NotFound"
import Memo from "../page/Memo"

import {createBrowserRouter, createHashRouter} from "react-router-dom";

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
                // path: '/board',
                index: true,    // 设置默认的二级路由
                element: <Board/>
            },
            {
                path: '/memo',
                element: <Memo/>
            }
        ]
    },
    {
        path: '*',
        element: <NotFound />
    }
])

export default router