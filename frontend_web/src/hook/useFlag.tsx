import React from 'react'
import { useFlagProps } from './types'

const useFlag = (): useFlagProps.returnProps => {
    return {
        flagToDom: ({
            flag, flagYText, flagNText
        }) => {
            const isRedFlag = flag === 'Y'

            return (
                <span
                    className={
                        isRedFlag
                            ? 'text-danger'
                            : 'text-success'
                    }
                >
                    {
                        isRedFlag
                            ? flagYText
                            : flagNText
                    }
                </span>
            )
        }
    }
}
export default useFlag
