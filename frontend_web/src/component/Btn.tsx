/* eslint-disable react/display-name */
import { DeleteOutlined, EditOutlined, SaveOutlined } from '@ant-design/icons'
import { Button } from 'antd'
import React from 'react'

const Btn = (): JSX.Element => {
    return (
        <>
        </>
    )
}
Btn.Edit = ({ onClick }: { onClick: () => void }): JSX.Element => {
    return (
        <>
            <Button
                style={{ color: 'blue' }}
                onClick={onClick}
                className={styles.editLabel}
                icon={<EditOutlined style={{ marginRight: '10px' }} />}
            >
                編輯
            </Button>
        </>
    )
}
Btn.Delete = ({ onClick }: { onClick: () => void }): JSX.Element => {
    return (
        <>
            <Button
                onClick={onClick}
                className={styles.editLabel}
                style={{ color: 'red' }}
                icon={<DeleteOutlined style={{ marginRight: '10px' }} />}
            >
                刪除
            </Button>
        </>
    )
}
Btn.Save = ({ onClick }: { onClick: () => void }): JSX.Element => {
    return (
        <>
            <Button
                onClick={onClick}
                className={styles.editLabel}
                style={{ color: 'green' }}
                icon={<SaveOutlined style={{ marginRight: '10px' }} />}
            >
                儲存
            </Button>
        </>
    )
}
Btn.Cancel = ({ onClick }: { onClick: () => void }): JSX.Element => {
    return (
        <>
            <Button
                onClick={onClick}
                className={styles.editLabel}
                style={{ color: 'skyblue' }}
                icon={<SaveOutlined style={{ marginRight: '10px' }} />}
            >
                取消
            </Button>
        </>
    )
}
export default Btn
