/* eslint-disable @typescript-eslint/strict-boolean-expressions */
import React from 'react'
import { useParams } from 'react-router-dom'
import Loading from '../../component/Loading'
const ShiftPage = (): JSX.Element => {
    const { banch } = useParams()
    if (!banch) {
        return (
            <Loading />
        )
    }
    return (
        <>
            {
                banch && (banch)
            }
            edit
        </>
    )
}
export default ShiftPage
