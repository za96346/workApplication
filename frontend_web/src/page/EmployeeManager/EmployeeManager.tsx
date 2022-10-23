import { SearchOutlined } from '@ant-design/icons'
import { DatePicker, Input, Spin } from 'antd'
import moment from 'moment'
import React, { useEffect, useRef, useState } from 'react'
import api from '../../api/api'
import BanchSelector from '../../component/BanchSelector'
import Btn from '../../component/Btn'
import PermessionSelector from '../../component/PermessionSelector'
import StatusSelector from '../../component/StatusSelector'
import useReduceing from '../../Hook/useReducing'
import statics from '../../statics'
import { UserType } from '../../type'

const editInit = {
    currentIdx: -1
}
const EmployeeManager = (): JSX.Element => {
    const { company, loading } = useReduceing()
    const form = useRef<UserType>({
        UserId: -1,
        UserName: '',
        CompanyCode: '',
        EmployeeNumber: '',
        OnWorkDay: '',
        Banch: -1,
        Permession: 2,
        WorkState: 'on'
    })
    const [edit, setEdit] = useState(editInit)
    const onSave = async (): Promise<void> => {
        console.log(form.current)
        const res = await api.updateUser(form.current)
        if (res.status) {
            await api.getUserAll()
        }
        setEdit((prev) => ({ ...prev, currentIdx: -1 }))
    }
    useEffect(() => {
        api.getUserAll()
    }, [])
    useEffect(() => {
        const user = company.employee?.find((item) => item?.UserId === edit.currentIdx)
        form.current = {
            ...user
        }
    }, [edit.currentIdx])
    return (
        <>
            <div className={styles.empManagerFilter}>
                <Input style={{ width: '150px' }} prefix={<SearchOutlined/>} placeholder='姓名' />
            </div>
            <div className={styles.empManagerTable}>
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
                            company.employee?.map((item) => {
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
                                                    defaultValue={moment(item.OnWorkDay)}
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
                                if (
                                    loading.onUpdateUser) {
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
