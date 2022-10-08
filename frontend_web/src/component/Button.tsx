import React from 'react'

interface props {
    text: string
    onClick: any
}

export const Button = ({ text, onClick }: props): JSX.Element => {
    return (
        <>
            <button
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
