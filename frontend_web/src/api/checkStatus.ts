import { FullMessage } from '../method/notice'

class checkStatus {
    a: string
    constructor () {
        this.a = ''
    }

    static async Login (code: number): Promise<any> {
        if (code === 400) {
            await FullMessage.error(window.language.accountOrPasswordError)
        }
        if (code === 401) {
            await FullMessage.error(window.language.noUser)
        }
        if (code === 417) {
            await FullMessage.error(window.language.formatError)
        }
        if (code === 200) {
            await FullMessage.success(window.language.loginSuccess)
        }
    }
}
export default checkStatus
