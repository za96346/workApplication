import { Tabs } from 'antd'
import React, { useEffect } from 'react'
import api from 'api/api'
import { WaitReplyType } from 'type'
import PeopleList from './PeopleList'

const SignPage = (): JSX.Element => {
    const onConfirm = async (WaitId: WaitReplyType['WaitId'], isAccept: WaitReplyType['IsAccept']): Promise<void> => {
        const res = await api.updateWaitReply({
            SpecifyTag: '',
            IsAccept: isAccept,
            WaitId
        })
        if (res.status) {
            await api.getWaitReply()
        }
    }
    useEffect(() => {
        api.getWaitReply()
    }, [])
    return (

        <div className={window.styles.signBlock}>
            <Tabs>
                <Tabs.TabPane tab={'申請中'} key={0}>
                    <PeopleList keys={1} onConfirm={onConfirm}/>
                </Tabs.TabPane>
                <Tabs.TabPane tab={'拒絕'} key={1}>
                    <PeopleList keys={3} onConfirm={onConfirm}/>
                </Tabs.TabPane>
                <Tabs.TabPane tab={'接受'} key={2}>
                    <PeopleList keys={2} onConfirm={onConfirm}/>
                </Tabs.TabPane>
            </Tabs>
        </div>
    )
}
export default SignPage
