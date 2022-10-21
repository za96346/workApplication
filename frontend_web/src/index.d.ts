declare module '*.module.css' {
    const classes: { readonly [ key: string ]: string }
    export default classes
}

declare module '*.module.sass' {
    const classes: { readonly [ key: string ]: string }
    export default classes
}

declare module '*.module.scss' {
    const classes: { readonly [ key: string ]: string }
    export default classes
}
declare interface Window {
    styles: object
}
declare module '*.webp' {
    const classes: string
    export default classes
}
declare module '*.png' {
    const classes: string
    export default classes
}
declare const styles: any
