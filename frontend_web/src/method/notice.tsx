import { message, notification } from 'antd'

export const openNotification = (): void => {
    notification.open({
        message: 'Notification Title',
        description:
      'This is the content of the notification. This is the content of the notification. This is the content of the notification.',
        onClick: () => {
            console.log('Notification Clicked!')
        }
    })
}

export const FullMessage = (): void => {}
FullMessage.success = (text: string) => message.success(text)
FullMessage.error = (text: string) => message.error(text)
