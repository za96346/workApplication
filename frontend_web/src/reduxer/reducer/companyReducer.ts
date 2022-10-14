import { action, companyType } from "../types"

const companyState = {
    banch: []
}

export const companyReducer = (state = companyState, action: action): any => {
    switch (action.type) {
        case companyType.SET_BANCH:
            return {
                banch: action.payload.banch
            }
        default:
            return state
    }
}
