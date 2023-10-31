import React from 'react'
import { Descriptions, Spin } from 'antd'

import useReduceing from 'Hook/useReducing'
// import { Button as MainBtn } from '../../../../component/Button'

import statics from '../../../../statics'
import dateHandle from 'method/dateHandle'

const PersonalDescibe = (): JSX.Element => {
    const { loading, user } = useReduceing()

    if (loading.onFetchSelfData) {
        return <Spin size='large' />
    }
    return (
        <>
            <Descriptions className='w-100' title="個人資訊" bordered>
                <Descriptions.Item label="員工編號" span={2}>{user.selfData?.EmployeeNumber}</Descriptions.Item>
                <Descriptions.Item label="帳號" span={2}>{user.selfData.Account}</Descriptions.Item>
                <Descriptions.Item label="姓名" span={2}>{user.selfData.UserName}</Descriptions.Item>
                <Descriptions.Item label="公司編號" span={2}>{user.selfData.CompanyCode}</Descriptions.Item>
                <Descriptions.Item label="到職日" span={2}>{`${dateHandle.transferUtcFormat(user.selfData?.OnWorkDay)}`}</Descriptions.Item>
                <Descriptions.Item label="部門" span={2}>{user.selfData?.BanchName}</Descriptions.Item>
                <Descriptions.Item label="職位" span={2}>
                    {statics.permession[user.selfData.Permession]}
                </Descriptions.Item>
            </Descriptions>
            {/* <MainBtn
                text='更改密碼'
                onClick={() => setStatus((prev) => ({ ...prev, changePwdBtn: true }))}
            /> */}

        </>
    )
}
export default PersonalDescibe
