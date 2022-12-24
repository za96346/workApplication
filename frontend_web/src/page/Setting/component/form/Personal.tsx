import React from "react"

const PersonalForm = (): JSX.Element => {
    return (
        <>
            {/* <Form
                name="basic"
                initialValues={{ remember: true }}
                onFinish={async (v) => await onFinish(v, 1)}
                autoComplete="off"
            >
                <Form.Item
                    label="帳號"
                    name="Account"
                    initialValue={user.selfData?.Account || ''}
                >
                    <Input disabled />
                </Form.Item>
                <Form.Item
                    label="姓名"
                    name="UserName"
                    initialValue={user.selfData?.UserName || ''}
                >
                    <Input />
                </Form.Item>
                {
                    user.selfData?.CompanyCode !== '' && (
                        <><Form.Item
                            label="公司編號"
                            name="CompanyCode"
                            initialValue={user.selfData?.CompanyCode || ''}
                        >
                            <Input disabled />
                        </Form.Item><Form.Item
                            label="到職日"
                            name="OnWorkDay"
                            initialValue={moment(user.selfData?.OnWorkDay || '2001-07-01')}
                        >
                            <DatePicker allowClear={false} disabled />
                        </Form.Item><Form.Item
                            label="部門"
                            name="Banch"
                        >
                            <BanchSelector disabled defaultValue={user.selfData?.Banch || 0} />
                        </Form.Item><Form.Item
                            label="權限"
                            name="Permession"
                        >
                            <PermessionSelector disabled defaultValue={user.selfData?.Permession} />
                        </Form.Item></>
                    )
                }
                <Form.Item
                    style={{ marginTop: '130px' }}
                    name="username"
                >
                    <Button
                        style={{ width: '100%', height: '50px' }}
                        htmlType="submit"
                    >
                        修改
                    </Button>
                </Form.Item>
            </Form> */}
        </>
    )
}
export default PersonalForm
