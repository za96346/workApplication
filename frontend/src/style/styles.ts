

export const styleLogin = {
    
    Input: {

    },
}
export const shadowWrapper = (shadowColor: string, shadowOffset:object, shadowOpacity:number, shadowRadius:number | string): object => {
    return{
        shadowColor: shadowColor,
        shadowOffset: shadowOffset,
        shadowOpacity: shadowOpacity,
        shadowRadius: shadowRadius,
    }
}