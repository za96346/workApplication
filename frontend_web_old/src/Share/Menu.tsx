import { AppstoreOutlined, CalendarOutlined, EditFilled, ExportOutlined, FieldTimeOutlined, GoldOutlined, HomeOutlined, IdcardOutlined, InsertRowRightOutlined, MenuFoldOutlined, MenuUnfoldOutlined, SettingOutlined, WalletOutlined } from '@ant-design/icons'
import { Button, MenuProps, Menu } from 'antd'
import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useBreakPoint } from '../Hook/useBreakPoint'
import useReduceing from '../Hook/useReducing'
import { clearAll } from '../reduxer/clearAll'
import { BanchType, UserType } from '../type'

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

const items = (banch: BanchType[], selfData: UserType): MenuItem[] => {
    const a = banch?.map((item) => getItem(item.BanchName, `a${item.Id}`)) || []
    const b = banch?.map((item) => getItem(item.BanchName, `b${item.Id}`)) || []
    const c = banch?.reduce((accr, item) => {
        if (selfData.Permession === 100) {
            accr.push(getItem(item.BanchName, `c${item.Id}`))
        } else if (selfData.Permession === 1 && selfData.Banch === item.Id) {
            accr.push(getItem(item.BanchName, `c${item.Id}`))
        }
        return accr
    }, [selfData.Permession === 100 && getItem('主管', 'c-200')])
    if (selfData.CompanyCode === '') {
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

        getItem('班表設定', 'shiftSetting', <AppstoreOutlined />, b),
        getItem('績效評核', 'performance', <WalletOutlined />, c.length > 0 && c),
        getItem('年度績效', 'yearPerformance', <CalendarOutlined />),
        getItem('員工管理', 'employeeManager', <IdcardOutlined />),
        getItem('部門管理', 'banchManager', <GoldOutlined />),
        getItem('時數管理', 'workTimeManager', <FieldTimeOutlined />),
        // getItem('平假日設定', 'weekendSetting', <CalendarOutlined/>),
        getItem('申請中', 'sign', <EditFilled />),
        getItem('設定', 'setting', <SettingOutlined />, [
            getItem('個人資料', 'z1000'),
            getItem('公司資料', 'z1001')
        ])
    ])
}
const App: React.FC = () => {
    const { isMore, isLess } = useBreakPoint()
    const { loading, company, user } = useReduceing()
    const [current, setCurrent] = useState<any>({
        keyPath: 'shift',
        key: ''
    })
    const [collapsed, setCollapsed] = useState(!isLess('md'))
    const navigate = useNavigate()

    const onClick: MenuProps['onClick'] = e => {
        // console.log('click ', e)
        setCurrent(e)
    }

    const width = (): string | number => {
        if (isLess('md')) {
            return '100%'
        }
        return collapsed ? '50px' : 246
    }

    useEffect(() => {
        const [path1, path2] = current.keyPath
        // console.log(current)
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
            return
        }
        if (path1 === 'performance' || path2 === 'performance') {
            navigate(`performance/${current.key}`)
            return
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
        if (path1 === 'workTimeManager') {
            navigate('workTimeManager')
        }
        if (path1 === 'yearPerformance') {
            navigate('yearPerformance')
        }
        // if (path1 === 'weekendSetting') {
        //     navigate('weekendSetting')
        // }
    }, [current])

    return (
        <>

            <div style={{
                width: width(),
                transition: '0.2s'
            }}
            className={window.styles.menuBlock}
            >
                {
                    isMore('md') && (
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
                    )
                }
                {/* <div
                    style={{

                    }}
                > */}
                <Menu
                    disabled={loading.onFetchBanch}
                    inlineCollapsed={collapsed}
                    overflowedIndicator
                    theme={'light'}
                    onClick={onClick}
                    style={{
                        width: width(),
                        marginTop: isLess('md') ? '0px' : '30px',
                        height: isLess('md') ? '80vh' : '90vh',
                        maxHeight: '90vh',
                        overflow: 'scroll'
                    }}
                    defaultOpenKeys={['sub1']}
                    selectable
                    mode="inline"
                    items={items(company.banch, user.selfData)}
                />
                {/* </div> */}
                <div
                    style={{
                        width: width(),
                        transition: '0.2s',
                        padding: isLess('md') ? '10px 34px 10px 34px' : '0px'
                    }}
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