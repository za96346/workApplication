import { Space, Tag } from 'antd'
import React from 'react'
import userTypes from 'types/user'

interface props {
    selected: userTypes.TABLE[]
    setSelected: any
}

const UserTags = ({
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
                        {item?.UserName}
                    </Tag>
                ))
            }
        </Space>
    )
}
export default UserTags
