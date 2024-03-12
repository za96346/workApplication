import { DefaultOptionType } from 'antd/es/select'

export enum modalType {
    edit = 'edit',
    delete = 'delete',
    add = 'add',
    copy = 'copy',
    print = 'print',
    changeBanch = 'changeBanch',
    inquire = 'inquire'
}

export enum modalTitle {
    edit = '編輯',
    delete = '刪除',
    add = '新增',
    copy = '複製',
    print = '列印'
}

export enum quitWorkStatus {
    Y = '離職',
    N = '在職'
}

export enum ScopeNameEnum {
    all = '全部',
    self = '自己',
    customize = '自訂'
}

// 離職狀態下拉清單
export const quitWorkStatusSelectList: DefaultOptionType[] = [
    {
        value: 'Y',
        label: quitWorkStatus.Y
    },
    {
        value: 'N',
        label: quitWorkStatus.N
    }
]
