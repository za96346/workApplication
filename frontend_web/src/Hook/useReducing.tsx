import { useSelector } from 'react-redux'
import { useParams } from 'react-router-dom'
import { companyReducerType } from '../reduxer/reducer/companyReducer'
import { shiftReducerType } from '../reduxer/reducer/shiftEditReducer'
import { statusReducerType } from '../reduxer/reducer/statusReducer'
import { userReducerType } from '../reduxer/reducer/userReducer'
import { RootState } from '../reduxer/store'

interface props {
    company: companyReducerType
    user: userReducerType
    loading: statusReducerType
    shiftEdit: shiftReducerType
    _persist: any
    state: {
        banchId: number
        banchName: string
    }
}

const useReduceing = (): props => {
    const {
        company,
        user,
        status: loading,
        shiftEdit,
        _persist
    }: RootState = useSelector((state: RootState) => state)
    // 參數
    const { banchId: currentBanchId } = useParams()
    const banchId = parseInt(currentBanchId?.replace('c', '')?.replace('b', '')?.replace('a', ''))
    const banchName = company.banch?.find((item: { Id: number }) => item.Id === banchId)?.BanchName || ''
    return {
        company,
        user,
        loading,
        shiftEdit,
        _persist,
        state: {
            banchId, // 現在網頁參數的banchId
            banchName
        }
    }
}
export default useReduceing
