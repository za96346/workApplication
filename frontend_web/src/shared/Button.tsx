import { Button, Form } from 'antd'
import type { ButtonProps, FormInstance } from 'antd'
import React from 'react'

const Btn = {
    Submit: ({ form, text }: { form: FormInstance, text: string }) => {
        const [submittable, setSubmittable] = React.useState(false)

        // Watch all values
        const values = Form.useWatch([], form)

        React.useEffect(() => {
            form.validateFields().then(
                () => {
                    setSubmittable(true)
                },
                () => {
                    setSubmittable(false)
                }
            )
        }, [values])

        return (
            <Button type="primary" htmlType="submit" disabled={!submittable}>
                {text}
            </Button>
        )
    },
    Add: (attr: ButtonProps) => (
        <Button
            type="primary"
            htmlType="button"
            {...attr}
        >
            新增
        </Button>
    ),
    Save: (attr: ButtonProps) => (
        <Button
            type="primary"
            htmlType="submit"
            {...attr}
        >
            儲存
        </Button>
    ),
    Cancel: (attr: ButtonProps) => (
        <Button
            type="primary"
            htmlType="button"
            {...attr}
        >
            取消
        </Button>
    ),
    Close: (attr: ButtonProps) => (
        <Button
            type="primary"
            htmlType="button"
            {...attr}
        >
            關閉
        </Button>
    )
}

export default Btn
