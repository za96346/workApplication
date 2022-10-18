class statics {
    a: any
    constructor () {
        this.a = ''
    }

    personalSetting = '1000'
    companySetting = '1001'
    permession: {
        [index: number]: string
    } = {
            100: '管理員',
            1: '主管',
            2: '一般職員'
        }

    weekDay: {
        [index: number]: string
    } = {
            1: '星期ㄧ',
            2: '星期二',
            3: '星期三',
            4: '星期四',
            5: '星期五',
            6: '星期六',
            7: '星期日'
        }

    weekType: {
        [index: number]: string
    } = {
            1: '平日',
            2: '假日'
        }
}

export default new statics()
