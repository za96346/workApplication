import React from 'react'

interface props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    text: string
    onClick: any
    style?: any
}

export const Button = ({ text, onClick, style, ...attr }: props): JSX.Element => {
    const { className, ...other } = attr
    return (
        <>
            <button
                {...other}
                style={style}
                onClick={onClick}
                className={className?.length > 0 ? className : styles.mainBtn}
            >
                {
                    text
                }
            </button>
        </>
    )
}
