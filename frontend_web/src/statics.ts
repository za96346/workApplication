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

    workState = {
        off: '離職',
        on: '在職'
    }

    isAccept: {
        [index: number]: string
    } = {
            1: '等待確認',
            2: '接受',
            3: '拒絕'
        }

    days: {
        [index: number]: string
    } = {
            0: '日',
            1: '一',
            2: '二',
            3: '三',
            4: '四',
            5: '五',
            6: '六'
        }

    type = {
        edit: 'edit',
        create: 'create'
    }

    shiftSocketEvent = {
        position: 'position',
        shift: 'shift',
        done: 'done'
    }

    shiftSettingObj = {
        coEdit: '共同編輯',
        sortEdit: '順序編輯',
        assignEdit: '指定編輯'
    }

    scope = {
        global: '全部',
        self: '自己',
        exceptManage: '除主管外'
    }
}

export default new statics()
