import { FullMessage } from '../method/notice'
import language from '../language'

class checkStatus {
    a: string
    constructor () {
        this.a = ''
    }

    static async Login (code: number): Promise<any> {
        switch (code) {
            case 400:
                await FullMessage.error(language.accountOrPasswordError)
                break
            case 401:
                await FullMessage.error(language.noUser)
                break
            case 417:
                await FullMessage.error(language.formatError)
                break
            case 200:
                await FullMessage.success(language.loginSuccess)
                break
        }
    }

    static async GetEmailCaptcha (code: number): Promise<any> {
        if (code === 200) {
            await FullMessage.success(language.emailCaptchaSendedSuccess)
        }
        if (code === 404) {
            await FullMessage.error(language.emailCaptchaSendedFail)
        }
    }

    static async Register (code: number): Promise<any> {
        switch (code) {
            case 200:
                await FullMessage.success(language.registerSuccess)
                break
            case 409:
                await FullMessage.error(language.accountHasBeenRegistered)
                break
            case 403:
                await FullMessage.error(language.registerFail)
                break
            case 417:
                await FullMessage.error(language.formatError)
                break
            case 400:
                await FullMessage.error(language.captchaIsNotCorrect)
                break
            case 422:
                await FullMessage.error(language.passwordIsNotSame)
                break
            case 451:
                await FullMessage.error(language.emailIsNotCorrect)
                break
        }
    }
}
export default checkStatus
