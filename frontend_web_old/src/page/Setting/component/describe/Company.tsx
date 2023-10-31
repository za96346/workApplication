import React from 'react'
import { Descriptions, Spin } from 'antd'

import useReduceing from 'Hook/useReducing'
// import { Button as MainBtn } from '../../../../component/Button'

const CompanyDescibe = (): JSX.Element => {
    const { loading, company } = useReduceing()

    if (loading.onFetchSelfData) {
        return <Spin size='large' />
    }
    return (
        <>
            <Descriptions className='w-100' title="公司資訊" bordered>
                <Descriptions.Item label="公司碼">{company.info.CompanyCode}</Descriptions.Item>
                <Descriptions.Item label="公司名稱">{company.info.CompanyName}</Descriptions.Item>
                <Descriptions.Item label="結算日">{company.info.SettlementDate}</Descriptions.Item>
                <Descriptions.Item label="公司地址">{`${company.info.CompanyLocation}`}</Descriptions.Item>
                <Descriptions.Item label="公司電話" span={2}>{company.info.CompanyPhoneNumber}</Descriptions.Item>
            </Descriptions>
            {/* <MainBtn
                text='更改密碼'
                onClick={() => setStatus((prev) => ({ ...prev, changePwdBtn: true }))}
            /> */}

        </>
    )
}
export default CompanyDescibe
