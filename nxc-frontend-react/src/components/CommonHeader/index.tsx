import React from "react";
import styles from "./index.module.less"
import {Button, Flex, Drawer, theme} from "antd";
import useStore from "../../store"
// import {MenuOutline} from "@ricons/ionicons5"
// import Icon from "@ant-design/icons"
import CommonLogo from "../CommonLogo";
import AdminSider from "../../pages/Admin/AdminSider"

const CommonHeader: React.FC = () => {
    const {isMenuFolded, setIsShowMobileMenu, isShowMobileMenu, setTheme} = useStore();
    const toggleTheme = () => {
        setTheme(theme.darkAlgorithm)
    }
    return (
        <div className={styles.root}>
            <Flex className={styles.innerFlex} align={'center'} justify={'space-between'}>
                {isMenuFolded?(<Button
                        type={'primary'}
                        variant={'filled'}
                        onClick={() => setIsShowMobileMenu(true)}
                    >
                        显示菜单
                    </Button>):(<div>仪表盘</div>)
                }

                <div className={styles.innerRoot}>
                    <Button type={'text'} variant={'text'} onClick={() => toggleTheme() }>
                        切换主题
                    </Button>
                </div>
            </Flex>
            <Drawer
                placement={'left'}
                closable={false}
                onClose={() => setIsShowMobileMenu(false)}
                open={isShowMobileMenu}
                key={'mobileMenu'}
                width={'60%'}
                className={styles.drawerInner}
                style={{padding: 0}}
                rootClassName={styles.drawerInner}
            >
                <CommonLogo />
                <AdminSider/>
            </Drawer>

        </div>
    )
}

export default CommonHeader