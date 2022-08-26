import { actionType } from "../type/actionType";


const getDataState = {
    state: null,
    data: null
}

export const loginData = (state = getDataState, action) => {
    switch(action.type) {
        case actionType.PERSON_DATA_SUCCESS:
            return {
                state: true,
                data: action.payLoad
            };
        case actionType.PERSON_DATA_ERROR:
            return {
                state: false,
                data: action.payLoad
            };
        default:
            return state;
    }
};