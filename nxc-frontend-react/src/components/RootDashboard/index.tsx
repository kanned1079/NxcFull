import styles from "./index.module.less";
import React from "react";
import {Layout} from "antd";
import useStore from "../../store"
import CommonLogo from "../CommonLogo";
import AdminSider from "../../pages/Admin/AdminSider"
import CommonHeader from "../CommonHeader";

const {Header, Sider, Content} = Layout;


const RootDashboard: React.FC = () => {
    const {isMenuFolded} = useStore();
    return (
        <div className={styles.root} style={{height: '100vh', display: 'flex', flexDirection: 'column'}}>
            <Layout className={styles.rootLayout}>
                {!isMenuFolded && <Sider width={240} className={styles.innerSider}>
                        <div>
                            <CommonLogo/>
                            <AdminSider/>
                        </div>
                    </Sider>
                }
                <Layout className={styles.innerLayout}> {/* This ensures Content and Footer stretch */}
                    <Header className={styles.innerHeader}>
                        <CommonHeader/>
                    </Header>
                    <Content className={styles.innerContent}>Content</Content>
                    {/*<Footer style={footerStyle}>Footer</Footer>*/}
                </Layout>
            </Layout>
        </div>
    );
};

export default RootDashboard;