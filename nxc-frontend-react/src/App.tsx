import { ConfigProvider } from 'antd';
import RootDashboard from './components/RootDashboard';
import useStore from "./store"
import { useEffect } from "react";

const App = () => {
    const { glacierBlue, setIsMenuFolded, isMenuFolded } = useStore();

    useEffect(() => {
        const handleResize = () => {
            // 当页面宽度小于 768px 时自动折叠
            if (window.innerWidth < 768) {
                if (!isMenuFolded) {
                    setIsMenuFolded(true); // 折叠菜单
                }
            } else {
                if (isMenuFolded) {
                    setIsMenuFolded(false); // 展开菜单
                }
            }
        };

        // 初始检查，确保页面加载时根据当前窗口宽度设置菜单状态
        handleResize();

        // 添加窗口尺寸变化监听
        window.addEventListener("resize", handleResize);

        // 清理事件监听器
        return () => {
            window.removeEventListener("resize", handleResize);
        };
    }, [isMenuFolded, setIsMenuFolded]);

    return (
        <ConfigProvider theme={glacierBlue?.theme || {}}>
            <div className="App">
                <RootDashboard />
            </div>
        </ConfigProvider>
    );
};

export default App;