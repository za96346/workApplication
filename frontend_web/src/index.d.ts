declare module'*.scss' {
    const content: {[key: string]: any}
    export = content
}
// declare global {
//     interface Window { styles: any; }
//     var styles: any
// }
declare module '*.webp' {
    const webp: string
}
declare module '*.png' {
    const png: string
}