import {Outlet, Link} from "react-router-dom";

const Layout = () => {
    return (
        <div>
            <Link to={'/'}>转到Board</Link>
            <Link to={'/about'}>转到About</Link>
            <p>一级路由组件</p>
            <Outlet />
        </div>
    )
}

export default Layout
