import React from 'react'

interface props extends React.InputHTMLAttributes<HTMLInputElement> {
    title: string
}

const Input = ({
    title, ...attr
}: props): JSX.Element => {
    const { className, ...other } = attr
    return (
        <>
            <div className={className}>
                <div className="input-group">
                    <div className="input-group-prepend">
                        <span className="input-group-text" id="">{
                            title
                        }</span>
                    </div>
                    <input {...other} type="text" className="form-control" />
                </div>
            </div>
        </>
    )
}
export default Input
