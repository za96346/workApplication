import { AppstoreOutlined, ExportOutlined, IdcardOutlined, InsertRowRightOutlined, MenuFoldOutlined, MenuUnfoldOutlined, SettingOutlined } from '@ant-design/icons'
import { Button, MenuProps, Menu } from 'antd'
import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import userAction from '../reduxer/action/userAction'
import { RootState } from '../reduxer/store'
import { BanchType } from '../type'

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

// getItem('保育組', 2)

const items = (banch: BanchType[]): MenuItem[] => {
    const a = banch?.map((item) => getItem(item.BanchName, item.Id)) || []
    return ([
        getItem('排班', 'shift', <InsertRowRightOutlined />, a),

        getItem('班表設定', 'shiftSetting', <AppstoreOutlined />, a),
        getItem('員工管理', 'employeeManager', <IdcardOutlined />),

        getItem('設定', 'setting', <SettingOutlined />, [
            getItem('個人資料', 1000),
            getItem('公司資料', 1001)
        ])
    ])
}
const App: React.FC = () => {
    const dispatch = useDispatch()
    const { banch } = useSelector((state: RootState) => state.company)
    const { onFetchBanch } = useSelector((state: RootState) => state.status)
    const [current, setCurrent] = useState<any>({
        keyPath: 'shift',
        key: ''
    })
    const [collapsed, setCollapsed] = useState(true)
    const navigate = useNavigate()

    const onClick: MenuProps['onClick'] = e => {
        console.log('click ', e)
        setCurrent(e)
    }

    const width = (): string | number => {
        return collapsed ? '50px' : 246
    }

    useEffect(() => {
        const [path1, path2] = current.keyPath
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
                        disabled={onFetchBanch}
                        inlineCollapsed={collapsed}
                        overflowedIndicator
                        theme={'light'}
                        onClick={onClick}
                        style={{ width: width() }}
                        defaultOpenKeys={['sub1']}
                        selectedKeys={[current]}
                        mode="inline"
                        items={items(banch)}
                    />
                </div>
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
