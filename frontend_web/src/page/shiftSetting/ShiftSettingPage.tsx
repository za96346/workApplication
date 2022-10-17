import { DeleteOutlined, EditOutlined, PictureOutlined } from '@ant-design/icons'
import { Button, Collapse, Input, List, TimePicker, Form, Tabs } from 'antd'
import React from 'react'
import { useSelector } from 'react-redux'
import { useParams } from 'react-router-dom'
import api from '../../api/api'
import dateHandle from '../../method/dateHandle'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'
import { RootState } from '../../reduxer/store'
import { BanchStyleType, ShiftSettingListType } from '../../type'

const data = (arr: BanchStyleType[]): ShiftSettingListType[] => arr.map((item) => {
    return {
        id: item.StyleId,
        title: item.TimeRangeName,
        icons: item.Icon,
        time: <>
            <span>上班： {item.OnShiftTime}</span>
            <span style={{ marginLeft: '10px' }}>下班： {item.OffShiftTime}</span>
        </>
    }
})

const ShiftSettingPage = (): JSX.Element => {
    const { banchId } = useParams()
    const company: companyReducerType = useSelector((state: RootState) => state.company)
    const onFinish = async (v: any): Promise<void> => {
        console.log(v)
        const banchIdNumber = parseInt(banchId)
        const res = await api.createBanchStyle(
            {
                ...v,
                OnShiftTime: dateHandle.dateFormatToTime(v.OnShiftTime._d),
                OffShiftTime: dateHandle.dateFormatToTime(v.OffShiftTime._d),
                BanchId: banchIdNumber
            }
        )
        if (res.status) {
            await api.getBanchStyle(banchIdNumber)
        }
    }
    return (
        <>
            <Tabs>
                <Tabs.TabPane tab={'圖標設定'} key={0}>
                    <div className={styles.ShiftSettingEdit}>
                        <Collapse style={{ width: '100%' }} defaultActiveKey={['1']}>
                            <Collapse.Panel header="新增" key="1">

                                <Form scrollToFirstError onFinish={onFinish}>
                                    <Form.Item initialValue="" label="排班圖標" name="Icon">
                                        <Input style={{ marginBottom: '20px' }} placeholder='新增班別圖標' prefix={<PictureOutlined />} />
                                    </Form.Item>
                                    <Form.Item initialValue="" label="班別名稱" name="TimeRangeName">
                                        <Input style={{ marginBottom: '20px' }} placeholder='新增班別名稱' prefix={<EditOutlined />} />
                                    </Form.Item>
                                    <Form.Item initialValue="" label="上班時段" name="OnShiftTime">
                                        <TimePicker style={{ marginBottom: '20px', width: '100%' }} placeholder='新增上班時段' />
                                    </Form.Item>
                                    <Form.Item initialValue="" label="下班時段" name="OffShiftTime">
                                        <TimePicker style={{ marginBottom: '20px', width: '100%' }} placeholder='新增下班時段' />
                                    </Form.Item>
                                    <Form.Item>
                                        <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
                                            <Button htmlType='submit'>新增</Button>
                                        </div>
                                    </Form.Item>

                                </Form>

                            </Collapse.Panel>
                        </Collapse>
                    </div>
                    <div className={styles.shiftSettingBlock}>
                        <List
                            itemLayout="horizontal"
                            dataSource={data(company.banchStyle)}
                            renderItem={(item: ShiftSettingListType) => (
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
                </Tabs.TabPane>
                <Tabs.TabPane tab="排班規則" key={1}>

                </Tabs.TabPane>
            </Tabs>
        </>
    )
}
export default ShiftSettingPage
