import React from "react"
import { Icon } from "native-base"


export interface interfaceCO_InputProps {
    placeholder: string,
    icons: JSX.Element,
    change: Function,
    styIdx: string,
};
export interface interfaceCO_InputState{
    focusEvent: boolean,
    focusStyle: object
    blurStyle: object,
    isPress: boolean,
}
export interface interfaceLoginState {
    account: string,
    password: string,
    modalVisible: boolean,
    isLoading: boolean,
}
export type navigation = {
    navigation: any,
    route: any,
}

export interface interfaceCO_WarningWindowProps extends 
    typeCO_ButtonProps {
    //擴展
    visible: boolean,
    titleText:string,
    bodyText:string,
    navigation: Function,
}
export type typeCO_ButtonProps = {
    //元件按鈕的type
    styIdx: string,
    btnText: string, 
    btnAction: Function
}
export interface interfaceCO_LoadingProps {
    isVisible: boolean,
    styIdx: string
}
export interface interfaceCO_LoadingState {
    dot: string,
}
export interface interfaceCO_WorkListProps {
    styIdx: string,
    data: Array <{
        imgUrl: string,
        userName: string,
        todayWorkType: string,
        todayWorkTime: string
    }>,
}




export interface interfaceCO_NavigateListProps {
    icons: JSX.Element,
    name: string,
    style?: object,
    btnAction: Function,
    styIdex: string,
}

export interface interfaceCO_FormUIProps {
    // btnAction: Function,
    Label: string,
    imgUrl: any,
    btnAction: Function,
    styIdex: string,
}