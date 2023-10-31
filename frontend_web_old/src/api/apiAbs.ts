abstract class apiAbs {
    protected readonly route: any

    protected async GET (url: any, params: any, callBack: any): Promise<any> {}

    protected async POST (url: any, body: any, callBack: any, params?: any): Promise<any> {}

    protected async PUT (url: any, body: any, callBack: any, params?: any): Promise<any> {}

    protected async DELETE (url: any, body: any, callBack: any, Params?: any): Promise<any> {}
}

export default apiAbs
