import React from 'react'

interface props {
    text: string
    onClick: any
    style?: any
}

export const Button = ({ text, onClick, style }: props): JSX.Element => {
    return (
        <>
            <button
                style={style}
                onClick={onClick}
                className={styles.mainBtn}
            >
                {
                    text
                }
            </button>
        </>
    )
}
