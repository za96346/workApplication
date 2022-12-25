import { Modal } from "antd"
import React from "react"
import { performanceType } from "Root/type"

interface props {
    open: boolean
    value?: performanceType | null
    onClose: () => void
}

const ChangeBanch = ({ open, value, onClose }: props): JSX.Element => {
    return (
        <Modal
            title="更換部門"
            destroyOnClose
            open={open}
            onCancel={onClose}
            onOk={onClose}
            okText="儲存"
            cancelText="取消"
        >
            hi
        </Modal>
    )
}
export default ChangeBanch
