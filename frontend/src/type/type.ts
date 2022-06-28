import React from "react";

export interface interfaceComponentInput {
    placeholder: string,
    styles: object,
    change: Function,
    require: NodeRequire;
};
export type shadowProps={
    shadowColor:string,
    shadowOffset:{
        width:number|string,
        height:number|string
    },
    shadowOpacity:number,
    shadowRadius:number,
}

