import { AppstoreOutlined, CalendarOutlined, EditFilled, ExportOutlined, GoldOutlined, HomeOutlined, IdcardOutlined, InsertRowRightOutlined, MenuFoldOutlined, MenuUnfoldOutlined, SettingOutlined } from '@ant-design/icons'
import { Button, MenuProps, Menu } from 'antd'
import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import useReduceing from '../Hook/useReducing'
import { clearAll } from '../reduxer/clearAll'
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

const items = (banch: BanchType[], companyCode: string): MenuItem[] => {
    const a = banch?.map((item) => getItem(item.BanchName, item.Id)) || []
    if (companyCode === '') {
        return (
            [
                getItem('首頁', 'home', <HomeOutlined />),
                getItem('設定', 'setting', <SettingOutlined />, [
                    getItem('個人資料', 1000)
                ])
            ]
        )
    }
    return ([
        getItem('排班', 'shift', <InsertRowRightOutlined />, a),

        getItem('班表設定', 'shiftSetting', <AppstoreOutlined />, a),
        getItem('員工管理', 'employeeManager', <IdcardOutlined />),
        getItem('部門管理', 'banchManager', <GoldOutlined />),
        getItem('平假日設定', 'weekendSetting', <CalendarOutlined/>),
        getItem('申請中', 'sign', <EditFilled />),
        getItem('設定', 'setting', <SettingOutlined />, [
            getItem('個人資料', 1000),
            getItem('公司資料', 1001)
        ])
    ])
}
const App: React.FC = () => {
    const { loading, company, user } = useReduceing()
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
        if (path1 === 'banchManager') {
            navigate('banchManager')
        }
        if (path1 === 'home') {
            navigate('home')
        }
        if (path1 === 'sign') {
            navigate('sign')
        }
        if (path1 === 'weekendSetting') {
            navigate('weekendSetting')
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
                        disabled={loading.onFetchBanch}
                        inlineCollapsed={collapsed}
                        overflowedIndicator
                        theme={'light'}
                        onClick={onClick}
                        style={{ width: width() }}
                        defaultOpenKeys={['sub1']}
                        selectedKeys={[current]}
                        mode="inline"
                        items={items(company.banch, user.selfData?.CompanyCode)}
                    />
                </div>
                <div
                    style={{
                        width: width(),
                        transition: '0.2s'
                    }}
                    className={styles.logout}
                    onClick={() => {
                        clearAll()
                        navigate('/entry/login', { replace: true })
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
