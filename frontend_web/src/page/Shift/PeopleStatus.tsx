import { UserOutlined } from '@ant-design/icons'
import React from 'react'
import { Badge, Avatar } from 'antd'

interface PeopleStatusProps {
    currentStatus: 'online'
    name: string
    color: string
}

const PeopleStatus = ({ currentStatus, name, color }: PeopleStatusProps): JSX.Element => {
    return (
        <>
            <div className={styles.peopleList}>
                <Avatar
                    icon={<UserOutlined />}
                />
                <div className={styles.peopleText}>
                    <div>{name}</div>
                    <div>編輯顏色: {<Badge color={color} status="error" />}</div>
                </div>
                <Badge text={'上線中'} status="success" />
            </div>
        </>
    )
}
export default PeopleStatus
