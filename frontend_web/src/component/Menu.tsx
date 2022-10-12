import { AppstoreOutlined, ExportOutlined, IdcardOutlined, InsertRowRightOutlined, MenuFoldOutlined, MenuUnfoldOutlined, SettingOutlined } from '@ant-design/icons'
import { Button, MenuProps, MenuTheme, Menu } from 'antd'
import React, { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import userAction from '../reduxer/action/userAction'

type MenuItem = Required<MenuProps>['items'][number]

const getItem = (
    label: React.ReactNode,
    key?: React.Key | null,
    icon?: React.ReactNode,
    children?: MenuItem[],
    type?: 'group'
): MenuItem => (
    {
        key,
        icon,
        children,
        label,
        type
    } as MenuItem)

const items: MenuItem[] = [
    getItem('排班', 'shift', <InsertRowRightOutlined />, [
        getItem('保育組', 2),
        getItem('社工組', 3),
        getItem('行政組', 4),
        getItem('公關組', 5)
    ]),

    getItem('班表設定', 'shiftSetting', <AppstoreOutlined />, [
        getItem('保育組', 6),
        getItem('社工組', 7),
        getItem('行政組', 8),
        getItem('公關組', 9),
        getItem('保育組', 11),
        getItem('社工組', 23),
        getItem('行政組', 33),
        getItem('公關組', 94)
    // getItem('Submenu', 'sub3', null, [getItem('Option 7', '7'), getItem('Option 8', '8')]),
    ]),
    getItem('員工管理', 'employeeManager', <IdcardOutlined />),

    getItem('設定', 'setting', <SettingOutlined />, [
        getItem('個人資料', 1000),
        getItem('公司資料', 1001)
    ])
]
const App: React.FC = () => {
    const [theme, setTheme] = useState<MenuTheme>('light')
    const dispatch = useDispatch()
    const [current, setCurrent] = useState<any>({
        keyPath: 'shift',
        key: ''
    })
    const [collapsed, setCollapsed] = useState(true)
    const navigate = useNavigate()

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const changeTheme = (value: boolean): void => {
        setTheme(value ? 'dark' : 'light')
    }

    const onClick: MenuProps['onClick'] = e => {
        console.log('click ', e)
        setCurrent(e)
    }

    const width = (): string | number => {
        return collapsed ? '50px' : 246
    }

    useEffect(() => {
        const [path1, path2] = current.keyPath
        console.log(path1, path2)
        if (path1 === 'setting' || path2 === 'setting') {
            navigate(`setting/${current.key}`)
            return
        }
        if (path1 === 'employeeManager' || path2 === 'employeeManager') {
            navigate('employeeManager')
            return
        }
        if (path1 === 'shift' || path2 === 'shift') {
            navigate(`shift/${current.key}`)
            return
        }
        if (path1 === 'shiftSetting' || path2 === 'shiftSetting') {
            navigate(`shiftSetting/${current.key}`)
        }
    }, [current])

    return (
        <>

            <div style={{
                width: width(),
                transition: '0.2s'
            }}
            className={styles.menuBlock}
            >
                <Button
                    type="primary"
                    onClick={() => setCollapsed((prev) => !prev)}
                    style={{
                        transition: '0.2s',
                        width: width(),
                        marginBottom: 16,
                        position: 'absolute'
                    }}
                >
                    {collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                </Button>
                <div
                    style={{
                        marginTop: '30px',
                        height: '90vh',
                        maxHeight: '90vh',
                        overflow: 'scroll'
                    }}
                >
                    <Menu
                        inlineCollapsed={collapsed}
                        overflowedIndicator
                        theme={theme}
                        onClick={onClick}
                        style={{ width: width() }}
                        defaultOpenKeys={['sub1']}
                        selectedKeys={[current]}
                        mode="inline"
                        items={items}
                    />
                </div>
                {/* <Switch
                checked={theme === 'dark'}
                onChange={changeTheme}
                checkedChildren="Dark"
                unCheckedChildren="Light"
                style={{ position: 'absolute', bottom: '10px', left: '50%', transform: 'translateX(-50%)' }}
            /> */}
                <div
                    style={{
                        width: width(),
                        transition: '0.2s'
                    }}
                    className={styles.logout}
                    onClick={() => {
                        navigate('/entry/login', { replace: true })
                        dispatch(userAction.clearToken())
                    }}
                >
                    {
                        width() !== '50px' && ('登出')
                    }
                    <ExportOutlined size={30} />
                </div>

            </div>
        </>
    )
}

export default App
