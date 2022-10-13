import { DeleteOutlined, EditOutlined, PictureOutlined } from '@ant-design/icons'
import { Button, Collapse, Input, List, TimePicker } from 'antd'
import React from 'react'

const data = [
    {
        title: '平日早班',
        icons: '*',
        time: <>
            <span>上班：08:00</span>
            <span style={{ marginLeft: '10px' }}>下班：17:00</span>
        </>
    },
    {
        title: '平日中班',
        icons: '●',
        time: <>
            <span>上班：12:00</span>
            <span style={{ marginLeft: '10px' }}>下班：21:00</span>
        </>
    },
    {
        title: '平日晚班',
        icons: '❂',
        time: <>
            <span>上班：17:00</span>
            <span style={{ marginLeft: '10px' }}>下班：08:00</span>
        </>
    },
    {
        title: '假日早班',
        icons: '♡',
        time: <>
            <span>上班：17:00</span>
            <span style={{ marginLeft: '10px' }}>下班：08:00</span>
        </>
    },
    {
        title: '假日中班',
        icons: '☹',
        time: <>
            <span>上班：17:00</span>
            <span style={{ marginLeft: '10px' }}>下班：08:00</span>
        </>
    },
    {
        title: '假日晚班',
        icons: '♚',
        time: <>
            <span>上班：17:00</span>
            <span style={{ marginLeft: '10px' }}>下班：08:00</span>
        </>
    }
]

const ShiftSettingPage = (): JSX.Element => {
    return (
        <>
            <div className={styles.ShiftSettingEdit}>
                <Collapse style={{ width: '100%' }} defaultActiveKey={['1']}>
                    <Collapse.Panel header="新增" key="1">
                        <>
                            <Input style={{ marginBottom: '20px' }} placeholder='新增班別圖標' prefix={<PictureOutlined />} />
                            <Input style={{ marginBottom: '20px' }} placeholder='新增班別名稱' prefix={<EditOutlined />} />
                            <TimePicker style={{ marginBottom: '20px', width: '100%' }} placeholder='新增上班時段' />
                            <TimePicker style={{ marginBottom: '20px', width: '100%' }} placeholder='新增下班時段' />
                            <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
                                <Button>新增</Button>
                            </div>
                        </>
                    </Collapse.Panel>
                </Collapse>
            </div>
            <div className={styles.shiftSettingBlock}>
                <List
                    itemLayout="horizontal"
                    dataSource={data}
                    renderItem={item => (
                        <List.Item>
                            <List.Item.Meta
                                avatar={<span style={{ fontSize: '2rem' }} >{item.icons}</span>}
                                title={<a href="#">{item.title}</a>}
                                description={item.time}
                            />
                            <div className={styles.editLabel} style={{ color: 'blue' }}>
                                <EditOutlined style={{ marginRight: '10px' }} />
                                編輯
                            </div>
                            <div className={styles.editLabel} style={{ marginLeft: '20px', color: 'red' }}>
                                <DeleteOutlined style={{ marginRight: '10px' }} />
                                刪除
                            </div>
                        </List.Item>
                    )}
                />
            </div>
        </>
    )
}
export default ShiftSettingPage
