import Swal, { type SweetAlertOptions, type SweetAlertResult } from 'sweetalert2'
import withReactContent, { type ReactSweetAlert } from 'sweetalert2-react-content'

class MySwalAbs {
    public MySwal: typeof Swal & ReactSweetAlert

    // 這個是 要傳 redux 的包裝funcion
    public wrapChildren: (v: any) => JSX.Element

    // 這是能被套建轉換的 key
    private readonly canBeTransToJSXKey: string[]

    constructor () {
        this.MySwal = withReactContent(Swal)
        this.canBeTransToJSXKey = [
            'title',
            'html',
            'confirmButtonText',
            'denyButtonText',
            'cancelButtonText',
            'footer',
            'closeButtonHtml',
            'iconHtml',
            'loaderHtml'
        ]
    }

    public WrapIntoRedux (value: SweetAlertOptions<any, any>):
    Record<string, any> {
        Object.keys(value).forEach((key) => {
            value[key] = this.canBeTransToJSXKey.includes(key)
                ? this.wrapChildren(value[key])
                : value[key]
        })
        return value
    }

    public async AlertMessage<T>(
        o: SweetAlertOptions<any, any>
    ): Promise<SweetAlertResult<T>> {
        return this.MySwal.fire({
            icon: 'warning',
            customClass: o.customClass,
            confirmButtonText: '確認',
            allowOutsideClick: true,
            ...this.WrapIntoRedux(o || {})
        }).then((e): any => {
            if (!e.isConfirmed) {
                return Promise.reject(e)
            }
            return Promise.resolve(e)
        })
    }

    public async checkMessage<T>(
        o: SweetAlertOptions<any, any>
    ): Promise<SweetAlertResult<T>> {
        return this.MySwal.fire({
            icon: 'question',
            showLoaderOnConfirm: true,
            showCancelButton: true,
            reverseButtons: true,
            allowOutsideClick: false,
            confirmButtonText: o?.confirmButtonText || '確認',
            cancelButtonText: o?.cancelButtonText || '取消',
            imageWidth: o?.imageUrl ? 400 : '',
            imageHeight: o?.imageUrl ? 200 : '',
            ...this.WrapIntoRedux(o || {})
        }).then((e): any => {
            if (!e.isConfirmed) {
                return Promise.reject(e)
            }
            return Promise.resolve(e)
        })
    }
}

const MySwal = new MySwalAbs()

export {
    MySwal,
    SweetAlertOptions as MySwalOptions
}
