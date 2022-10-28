import { useSelector } from "react-redux"
import { companyReducerType } from "../reduxer/reducer/companyReducer"
import { shiftReducerType } from "../reduxer/reducer/shiftEditReducer"
import { statusReducerType } from "../reduxer/reducer/statusReducer"
import { userReducerType } from "../reduxer/reducer/userReducer"
import { RootState } from "../reduxer/store"

interface props {
    company: companyReducerType
    user: userReducerType
    loading: statusReducerType
    shiftEdit: shiftReducerType
    _persist: any
}

const useReduceing = (): props => {
    const { company, user, status: loading, shiftEdit, _persist }: RootState = useSelector((state: RootState) => state)
    return {
        company, user, loading, shiftEdit, _persist
    }
}
export default useReduceing
