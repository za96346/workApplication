import React, { useEffect } from 'react'
import ShiftTable from '../ShiftTable'

const PrintShift = (): JSX.Element => {
    useEffect(() => {
        window.print()
    }, [])
    return (
        <div
            translate='no'
            className={window.styles.print_page}
            style={{
                transform: 'rotate(90deg)'
            }}
        >
            <ShiftTable />
        </div>
    )
}
export default PrintShift
