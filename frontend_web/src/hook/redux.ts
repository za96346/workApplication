import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux'
import { RootState, AppDispatch, allAction } from 'reducer/store'

// Use throughout your app instead of plain `useDispatch` and `useSelector`
type DispatchFunc = () => {
    dispatch: AppDispatch
    action: typeof allAction
}
export const useAppDispatch: DispatchFunc = () => {
    return {
        dispatch: useDispatch(),
        action: allAction
    }
}
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector
