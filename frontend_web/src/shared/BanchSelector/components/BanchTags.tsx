import { Space, Tag } from 'antd'
import React from 'react'
import companyBanchTypes from 'types/companyBanch'

interface props {
    selected: companyBanchTypes.TABLE[]
    setSelected: any
}

const BanchTags = ({ selected, setSelected }: props): JSX.Element => {
    return (
        <Space wrap>
            {
                selected?.map((item) => (
                    <Tag
                        key={item?.BanchId}
                        onClose={() => {
                            setSelected((prev) => prev.filter((i) => i?.BanchId !== item?.BanchId))
                        }}
                        closable
                    >
                        {item?.BanchName}
                    </Tag>
                ))
            }
        </Space>
    )
}
export default BanchTags
