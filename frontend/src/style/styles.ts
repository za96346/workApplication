import { shadowProps } from "../type/type"

export const styleLogin={
    
    Input:{
        width: '100%',
        backgroundColor: '#rgb(35,230,55)',
        height: 60,
        borderRadius: 10,
        paddingLeft: 60,
    },
}
export const shadowWrapper=(shadowColor:string,shadowOffset:object,shadowOpacity:number,shadowRadius:number|string):object => {
    return{
        shadowColor:shadowColor,
        shadowOffset:shadowOffset,
        shadowOpacity:shadowOpacity,
        shadowRadius:shadowRadius,
    }
}