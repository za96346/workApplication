import { Avatar, Button, List } from 'antd'
import React from 'react'
import useReduceing from 'Hook/useReducing'
import statics from 'statics'
import { WaitReplyType } from 'type'
interface props {
    onConfirm: (a: WaitReplyType['WaitId'], isAccept: WaitReplyType['IsAccept']) => Promise<void>
    keys: WaitReplyType['IsAccept']
}
const PeopleList = ({ onConfirm, keys }: props): JSX.Element => {
    const { company } = useReduceing()
    const waitReply = company.waitReply?.filter((item) => item?.IsAccept === keys)
    return (
        <List
            itemLayout="horizontal"
            dataSource={waitReply}
            renderItem={item => (
                <List.Item style={{ flexWrap: 'wrap', position: 'relative' }}>
                    <List.Item.Meta
                        style={{ minWidth: '200px', marginBottom: '30px' }}
                        avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                        title={<>{item.UserName}</>}
                        description={<>{statics.isAccept[item?.IsAccept || 1]}</>}
                    />
                    <div style={{ position: 'absolute', right: '0px', bottom: '5px' }}>
                        <Button
                            onClick={async () => await onConfirm(item.WaitId, 2)}
                            style={{ backgroundColor: 'skyblue', color: 'white' }}
                        >
                                確認
                        </Button>
                        <Button onClick={async () => await onConfirm(item.WaitId, 3)}>拒絕</Button>
                    </div>
                </List.Item>
            )}
        />
    )
}
export default PeopleList
