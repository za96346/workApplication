import Insert from 'Share/Insert'
import { Button, Form } from 'antd'
import React from 'react'
import useReduceing from 'Hook/useReducing'
import api from 'api/api'

const PasswordForm = (): JSX.Element => {
    const { user } = useReduceing()
    const onFinish = async (v: any, types: 1 | 2): Promise<void> => {
        console.log(v)
        if (types === 1) {
            const res = await api.UpdateSelfData(v.UserName)
            console.log(res)
        } else if (types === 2) {
            const res = await api.changePassword({ ...v, Captcha: parseInt(v.Captcha) })
            console.log(res)
        }
    }
    return (
        <>
            <Form onFinish={async (v) => await onFinish(v, 2)} style={{ marginTop: '20px' }}>
                <Insert.Captcha email={user.selfData.Account} />
                <Insert.OldPwd />
                <Insert.Pwd textNum={2} />
                <Insert.PwdConfirm />
                <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                    <Button onClick={() => {}}>取消</Button>
                    <Button htmlType='submit'>送出</Button>
                </div>
            </Form>
        </>
    )
}
export default PasswordForm
