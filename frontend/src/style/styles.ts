

export const styleLogin = {
    
    Input: {

    },
}
export const shadowWrapper = (
    shadowColor?: string,
    shadowOffset?: object, 
    shadowOpacity?: number,
    shadowRadius?: number | string): object => {
    return{
        shadowColor: shadowColor || '#333',
        shadowOffset: shadowOffset || {width:10, height:10},
        shadowOpacity: shadowOpacity || 0.5,
        shadowRadius: shadowRadius || 10,
        elevation: 2,
    }
}