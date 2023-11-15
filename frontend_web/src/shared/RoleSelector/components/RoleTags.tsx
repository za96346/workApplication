import { Space, Tag } from 'antd'
import React from 'react'
import roleTypes from 'types/role'

interface props {
    selected: roleTypes.TABLE[]
    setSelected: any
}

const RoleTags = ({
    selected,
    setSelected
}: props): JSX.Element => {
    return (
        <Space wrap>
            {
                selected?.map((item) => (
                    <Tag
                        key={item?.RoleId}
                        onClose={() => {
                            setSelected((prev) => (
                                prev.filter((i) => i?.RoleId !== item?.RoleId)
                            ))
                        }}
                        closable
                    >
                        {item?.RoleName}
                    </Tag>
                ))
            }
        </Space>
    )
}
export default RoleTags
