abstract class api {
    protected route: any

    protected async GET (url: any, params: any, callBack: any): Promise<any> {}

    protected async POST (url: any, body: any, callBack: any, params?: any): Promise<any> {}

    protected async PUT (url: any, body: any, callBack: any, params?: any): Promise<any> {}
}

export default api
