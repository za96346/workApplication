import { Dispatch } from "react"
import { actionType } from "../type/actionType"

export const getLoginData = () => async (dispatch) => {
    console.log('it do')
    try {
        dispatch({
            type: actionType.PERSON_DATA_SUCCESS,
            payLoad: {
                id:0,
                name: 'john'
            }
        })
    }
    catch {
        dispatch({
            type: actionType.PERSON_DATA_SUCCESS,
            payLoad: 'error'
        })
    }
}