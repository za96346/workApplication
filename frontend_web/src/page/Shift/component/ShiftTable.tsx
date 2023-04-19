import React from 'react'
import Head from './tabEdit/Head'
import Row from './tabEdit/Row'

interface props {
    sendMessage?: (v: string) => void
}
const ShiftTable = ({
    sendMessage = () => {}
}: props): JSX.Element => {
    return (
        <>
            <table
                style={{ cursor: 'pointer' }}
                className='mb-5 table table-bordered table-hover table-striped'>
                <thead>
                    <Head/>
                </thead>
                <tbody>
                    <Row sendMessage={sendMessage} />
                </tbody>
            </table>
        </>
    )
}
export default ShiftTable
