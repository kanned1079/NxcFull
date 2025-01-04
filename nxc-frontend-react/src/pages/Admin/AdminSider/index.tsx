import React from "react";
import styles from "./index.module.less"

import { AppstoreOutlined, MailOutlined, SettingOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Menu } from 'antd';

type MenuItem = Required<MenuProps>['items'][number];

const items: MenuItem[] = [
    {
        key: 'key1',
        label: '仪表板',
        icon: <MailOutlined />,
    },
    {
        key: 'doc',
        label: '使用文档',
        icon: <MailOutlined />,
    },
    {
        key: 'sub1',
        label: '订阅',
        icon: <MailOutlined />,
        children: [
            {
                key: 'g1',
                label: '购买订阅',
                type: 'group',
            },
            {
                key: 'g2',
                label: '我的密钥',
                type: 'group',
            },
            {
                key: 'g3',
                label: '激活记录',
                type: 'group',
            },
            {
                key: 'g4',
                label: 'Item 2',
                type: 'group',
                children: [
                    { key: '3', label: 'Option 3' },
                    { key: '4', label: 'Option 4' },
                ],
            },
        ],
    },
    {
        key: 'sub2',
        label: 'Navigation Two',
        icon: <AppstoreOutlined />,
        children: [
            { key: '5', label: 'Option 5' },
            { key: '6', label: 'Option 6' },
            {
                key: 'sub3',
                label: 'Submenu',
                children: [
                    { key: '7', label: 'Option 7' },
                    { key: '8', label: 'Option 8' },
                ],
            },
        ],
    },
    {
        type: 'divider',
    },
    {
        key: 'sub4',
        label: 'Navigation Three',
        icon: <SettingOutlined />,
        children: [
            { key: '9', label: 'Option 9' },
            { key: '10', label: 'Option 10' },
            { key: '11', label: 'Option 11' },
            { key: '12', label: 'Option 12' },
        ],
    },
    {
        key: 'grp',
        label: 'Group',
        type: 'group',
        children: [
            { key: '13', label: 'Option 13' },
            { key: '14', label: 'Option 14' },
        ],
    },
];

const AdminSider: React.FC = () => {
    const onClick: MenuProps['onClick'] = (e) => {
        console.log('click ', e);
    };

    return (
        <div className={styles.root}>
            <Menu
                className={styles.menu}
                onClick={onClick}
                style={{ width: 'auto' }}
                defaultSelectedKeys={['1']}
                defaultOpenKeys={['sub1']}
                mode="inline"
                items={items}
                inlineIndent={12}
            />
        </div>
    )
}

export default AdminSider