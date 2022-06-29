import React from "react";

export interface interfaceComponentInput {
    placeholder: string,
    styles: object,
    change: Function,
    require: NodeRequire;
};

export interface interfaceLoginState {
    account:string,
    password:string,
    modalVisible:boolean,
}
export interface interfaceComponentModal {
    visible: boolean,
    titleText:string,
    bodyText:string,
    btnText:string,
    change:Function
}
export interface interfaceComponentInputState{
    focusEvent:boolean,
    focusStyle:object
    blurStyle:object
}
