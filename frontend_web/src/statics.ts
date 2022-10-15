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
}

export default new statics()
