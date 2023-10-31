import React from 'react'
import { Descriptions, Divider } from 'antd'
import { performanceType } from 'type'
import dateHandle from 'method/dateHandle'

const DescribeValue = ({ value }: { value: performanceType }): JSX.Element => (
    <>
        <Descriptions>
            <Descriptions.Item span={1} label="姓名">
                {value?.UserName || ''}
            </Descriptions.Item>
            <Descriptions.Item span={1} label="年度">
                {value?.Year || ''}
            </Descriptions.Item>
            <Descriptions.Item span={1} label="月份">
                {value?.Month || ''}
            </Descriptions.Item>
            <Descriptions.Item span={3} label="創建時間">
                {dateHandle.transferUtcFormat(value?.CreateTime) || ''}
            </Descriptions.Item>
            <Descriptions.Item span={3} label="最後修改">
                {dateHandle.transferUtcFormat(value?.LastModify) || ''}
            </Descriptions.Item>
            <Descriptions.Item span={3} label="狀態">
                {
                    value?.CompanyId === -1
                        ? <span className='text-danger'>離職</span>
                        : <span className='text-primary'>在職</span>
                }
            </Descriptions.Item>
        </Descriptions>
        <Divider />
    </>

)
export default DescribeValue
