import { Button, Form } from 'antd'
import type { FormInstance } from 'antd'
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
    }
}

export default Btn
