/* eslint-disable array-callback-return */
import { PlusCircleFilled, SearchOutlined } from '@ant-design/icons'
import { Button, DatePicker, Input, Spin } from 'antd'
import React, { useEffect, useMemo, useRef, useState } from 'react'
import api from '../../api/api'
import BanchSelector from '../../Share/BanchSelector'
import Btn from '../../Share/Btn'
import PermessionSelector from '../../Share/PermessionSelector'
import StatusSelector from '../../Share/StatusSelector'
import useReduceing from '../../Hook/useReducing'
import statics from '../../statics'
import { UserType } from '../../type'
import dayjs from 'dayjs'

const editInit = {
    currentIdx: -1
}
const filtInit = {
    userName: '',
    workState: 'on',
    banch: -1
}
const EmployeeManager = (): JSX.Element => {
    const { company, loading } = useReduceing()
    const [data, setData] = useState<UserType[]>([])
    const [filt, setFilt] = useState(filtInit)
    const companyFilter = useMemo(() => {
        // eslint-disable-next-line array-callback-return
        const res1 = data.filter((item) => {
            if (filt === filtInit) return item
            if (item.UserName !== filt.userName && filt.userName !== '') {
                return
            }
            if (item.WorkState !== filt.workState) {
                return
            }
            if (item.Banch !== filt.banch && filt.banch !== -1) {
                return
            }
            return item
        })
        return res1
    }, [filt, data])
    const form = useRef<UserType>({
        UserId: -1,
        UserName: '',
        CompanyCode: '',
        EmployeeNumber: '',
        OnWorkDay: '',
        Banch: -1,
        Permession: 2,
        WorkState: 'on',
        BanchName: '',
        CompanyName: '',
        CompanyId: -1
    })
    const [edit, setEdit] = useState(editInit)
    const onSave = async (): Promise<void> => {
        console.log(form.current)
        const res = await api.updateUser(form.current)
        if (res.status) {
            const idx = data.filter((item) => item.UserId !== form.current.UserId)
            setData([...idx, form.current])
        }
        setEdit((prev) => ({ ...prev, currentIdx: -1 }))
    }
    useEffect(() => {
        api.getUserAll()
    }, [])
    useEffect(() => {
        setData(company.employee)
    }, [company.employee])
    useEffect(() => {
        const user = data?.find((item) => item?.UserId === edit.currentIdx)
        form.current = {
            ...user
        }
    }, [edit.currentIdx])
    return (
        <>
            <div className={window.styles.empManagerFilter}>
                <Button onClick={() => setFilt(filtInit)} style={{ marginRight: '20px' }}>重設</Button>
                <Input
                    onChange={(e) => setFilt((prev) => ({ ...prev, userName: e.target?.value || '' }))}
                    style={{ width: '150px' }}
                    prefix={<SearchOutlined/>}
                    placeholder='姓名'
                />
                <BanchSelector onChange={(e) => setFilt((prev) => ({ ...prev, banch: e }))} defaultValue={0}/>
                <StatusSelector onChange={(e) => setFilt((prev) => ({ ...prev, workState: e }))} defaultValue={'on'}/>
                <button className='btn btn-secondary d-flex align-items-center m-1'>
                    <PlusCircleFilled style={{ marginRight: '5px' }}/>
                    新增員工
                </button>
            </div>
            <div className={window.styles.empManagerTable}>
                <table>
                    <thead>
                        <tr>
                            <td>員工編號</td>
                            <td>姓名</td>
                            <td>權限</td>
                            <td>部門</td>
                            <td>到職日</td>
                            <td>工作狀態</td>
                            <td></td>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            companyFilter?.map((item) => {
                                if (edit.currentIdx > -1 && edit.currentIdx === item.UserId) {
                                    return (
                                        <tr key={item.UserId}>
                                            <td>
                                                <Input onChange={(v) => { form.current.EmployeeNumber = v.target.value }} placeholder='請輸入員工編號' defaultValue={item.EmployeeNumber}/>
                                            </td>
                                            <td>{item.UserName}</td>
                                            <td>
                                                <PermessionSelector onChange={(v) => { form.current.Permession = v }} defaultValue={item.Permession}/>
                                            </td>
                                            <td>
                                                <BanchSelector onChange={(v) => { form.current.Banch = v }} defaultValue={item.Banch}/>
                                            </td>
                                            <td>
                                                <DatePicker
                                                    allowClear={false}
                                                    inputReadOnly
                                                    onChange={(v: any) => { form.current.OnWorkDay = v._d.toISOString() }}
                                                    defaultValue={dayjs(item.OnWorkDay)}
                                                />
                                            </td>
                                            <td>
                                                <StatusSelector onChange={(v) => { form.current.WorkState = v }} defaultValue={item.WorkState}/>
                                            </td>
                                            <td>
                                                <Btn.Save onClick={onSave} />
                                                <Btn.Cancel onClick={() => setEdit((prev) => ({ ...prev, currentIdx: -1 }))} />
                                            </td>
                                        </tr>
                                    )
                                }
                                if (loading.onUpdateUser) {
                                    <tr>
                                        <Spin/>
                                    </tr>
                                }
                                return (
                                    <tr key={item.UserId}>
                                        <td>{item.EmployeeNumber}</td>
                                        <td>{item.UserName}</td>
                                        <td>{statics.permession[item.Permession]}</td>
                                        <td>{company.banch?.find((items) => items?.Id === item?.Banch)?.BanchName || ''}</td>
                                        <td>{new Date(item?.OnWorkDay)?.toLocaleDateString()}</td>
                                        <td>{statics.workState[item.WorkState]}</td>
                                        <td>
                                            {
                                                edit.currentIdx === -1 && (
                                                    <Btn.Edit onClick={() => setEdit((prev) => ({ ...prev, currentIdx: item.UserId }))} />
                                                )
                                            }
                                        </td>
                                    </tr>
                                )
                            })
                        }
                    </tbody>
                </table>
            </div>
        </>
    )
}
export default EmployeeManager
