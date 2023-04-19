import { useMemo } from 'react'
import useReduceing from 'Hook/useReducing'
import dateHandle from 'method/dateHandle'

interface props {
    dayArray: string[]
}

const useShiftBasic = (): props => {
    const { shiftEdit } = useReduceing()
    // 日期
    const dayArray = useMemo(() => {
        const start = shiftEdit?.StartDay
        const end = shiftEdit?.EndDay
        const dayContainer = ['']
        for (let i = 0; i < new Date(end).getDate(); i++) {
            const res = dateHandle.addDays(start, i)
            dayContainer.push(res)
        }
        return [...dayContainer, '總時數']
    }, [shiftEdit?.StartDay, shiftEdit?.EndDay])

    return {
        dayArray
    }
}
export default useShiftBasic
