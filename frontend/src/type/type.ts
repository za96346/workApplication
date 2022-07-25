export interface interfaceComponentInputProps {
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
export interface interfaceComponentModalProps extends 
    typeComponentButtonProps {
    //擴展
    visible: boolean,
    titleText:string,
    bodyText:string,
    navigation: typeNavigation,
}
export interface interfaceComponentInputState{
    focusEvent:boolean,
    focusStyle:object
    blurStyle:object
}
export type typeComponentButtonProps = {
    //元件按鈕的type
    btnText: string, 
    change: Function
}
export type typeNavigation = {
    navigate: (pageName:string, params: object) => {},
    popToTop: ()=>{},
    goBack: ()=>{}
}