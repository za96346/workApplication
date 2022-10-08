import { Switch, Transfer } from 'antd'
import type { TransferDirection } from 'antd/es/transfer'
import React, { useState } from 'react'

interface RecordType {
    key: string
    title: string
    description: string
    disabled: boolean
}

const mockData: RecordType[] = Array.from({ length: 20 }).map((_, i) => ({
    key: i.toString(),
    title: `content${i + 1}`,
    description: `description of content${i + 1}`,
    disabled: i % 3 < 1
}))

const oriTargetKeys = mockData.filter(item => Number(item.key) % 3 > 1).map(item => item.key)

const TransferFrame: React.FC = () => {
    const [targetKeys, setTargetKeys] = useState<string[]>(oriTargetKeys)
    const [selectedKeys, setSelectedKeys] = useState<string[]>([])
    const [disabled, setDisabled] = useState(false)

    const handleChange = (
        newTargetKeys: string[],
        direction: TransferDirection,
        moveKeys: string[]
    ): any => {
        setTargetKeys(newTargetKeys)

        console.log('targetKeys: ', newTargetKeys)
        console.log('direction: ', direction)
        console.log('moveKeys: ', moveKeys)
    }

    const handleSelectChange = (sourceSelectedKeys: string[], targetSelectedKeys: string[]): any => {
        setSelectedKeys([...sourceSelectedKeys, ...targetSelectedKeys])

        console.log('sourceSelectedKeys: ', sourceSelectedKeys)
        console.log('targetSelectedKeys: ', targetSelectedKeys)
    }

    const handleScroll = (
        direction: TransferDirection,
        e: React.SyntheticEvent<HTMLUListElement, Event>
    ): any => {
        console.log('direction:', direction)
        console.log('target:', e.target)
    }

    const handleDisable = (checked: boolean): any => {
        setDisabled(checked)
    }

    return (
        <>
            <Transfer
                dataSource={mockData}
                titles={['Source', 'Target']}
                targetKeys={targetKeys}
                selectedKeys={selectedKeys}
                onChange={handleChange}
                onSelectChange={handleSelectChange}
                onScroll={handleScroll}
                render={item => item.title}
                disabled={disabled}
                oneWay
                style={{ marginBottom: 16 }}
            />
            <Switch
                unCheckedChildren="disabled"
                checkedChildren="disabled"
                checked={disabled}
                onChange={handleDisable}
            />
        </>
    )
}

export default TransferFrame
