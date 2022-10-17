import { DeleteOutlined, EditOutlined, PictureOutlined } from '@ant-design/icons'
import { Button, Collapse, Input, List, TimePicker } from 'antd'
import React from 'react'
import { useSelector } from 'react-redux'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'
import { RootState } from '../../reduxer/store'
import { BanchStyleType, ShiftSettingListType } from '../../type'

const data = (arr: BanchStyleType[]): ShiftSettingListType[] => arr.map((item) => {
    return {
        title: item.TimeRangeName,
        icons: item.Icon,
        time: <>
            <span>上班： {item.OnShiftTime}</span>
            <span style={{ marginLeft: '10px' }}>下班： {item.OffShiftTime}</span>
        </>
    }
})

const ShiftSettingPage = (): JSX.Element => {
    const company: companyReducerType = useSelector((state: RootState) => state.company)
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
                    dataSource={data(company.banchStyle)}
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
