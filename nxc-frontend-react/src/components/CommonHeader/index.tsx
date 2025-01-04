import React from "react";
import styles from "./index.module.less"
import {Button, Flex, Drawer} from "antd";
import useStore from "../../store"
// import {MenuOutline} from "@ricons/ionicons5"
// import Icon from "@ant-design/icons"
import AdminSider from "../../pages/Admin/AdminSider"

const CommonLogo: React.FC = () => {
    const {isMenuFolded, setIsShowMobileMenu, isShowMobileMenu} = useStore();
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

                <div>r</div>
            </Flex>
            <Drawer
                placement={'left'}
                closable={false}
                onClose={() => setIsShowMobileMenu(false)}
                open={isShowMobileMenu}
                key={'mobileMenu'}
                width={'60%'}
            >
                <AdminSider/>
            </Drawer>

        </div>
    )
}

export default CommonLogo