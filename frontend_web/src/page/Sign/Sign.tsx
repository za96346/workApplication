import { Avatar, Button, List } from 'antd'
import React from 'react'

const data = [
    {
        title: 'siou'
    },
    {
        title: 'siou'
    },
    {
        title: 'so'
    },
    {
        title: 'Ant'
    }
]
const SignPage = (): JSX.Element => {
    return (
        <div className={styles.signBlock}>
            <List
                itemLayout="horizontal"
                dataSource={data}
                renderItem={item => (
                    <List.Item style={{ flexWrap: 'wrap', position: 'relative' }}>
                        <List.Item.Meta
                            style={{ minWidth: '200px', marginBottom: '30px' }}
                            avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                            title={<a href="https://ant.design">{item.title}</a>}
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
