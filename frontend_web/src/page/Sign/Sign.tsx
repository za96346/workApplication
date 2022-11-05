import { Avatar, Button, List } from 'antd'
import React, { useEffect } from 'react'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'

const SignPage = (): JSX.Element => {
    const { company } = useReduceing()
    const waitReply = company.waitReply
    useEffect(() => {
        api.getWaitReply()
    }, [])
    return (
        <div className={styles.signBlock}>
            <List
                itemLayout="horizontal"
                dataSource={waitReply}
                renderItem={item => (
                    <List.Item style={{ flexWrap: 'wrap', position: 'relative' }}>
                        <List.Item.Meta
                            style={{ minWidth: '200px', marginBottom: '30px' }}
                            avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                            title={<a href="https://ant.design">{item.UserName}</a>}
                            description=""
                        />
                        <div style={{ position: 'absolute', right: '0px', bottom: '5px' }}>
                            <Button style={{ backgroundColor: 'skyblue', color: 'white' }}>確認</Button>
                            <Button>拒絕</Button>
                        </div>
                    </List.Item>
                )}
            />
        </div>
    )
}
export default SignPage
