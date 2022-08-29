import { lang } from "./type";

class language implements lang{
    constructor() {
        console.log('language is init');
    }

    setLang():any {
        // eslint-disable-next-line @typescript-eslint/no-unused-expressions
        const currentLanguage = (window.navigator.language)
        if (currentLanguage === 'zh-TW') {
            return this.ch
        } else {
            return this.en
        }
    }
    private ch() {

    }
    private en() {

    }
}
export default new language();