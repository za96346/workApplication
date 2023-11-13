/* eslint-disable prefer-promise-reject-errors */
import axios from 'axios'
import { type SweetAlertOptions } from 'sweetalert2'
// import { reducer } from '@vteam_components/redux'
import { store, allAction } from 'reducer/store'

// 這個 global 一定要引入 不然此api build 時候會讀不到 actionT and storeT
// import '@vteam_components/redux/global'
import { MySwal } from '../func/Swal'

// type
interface responseParams<T> {
    data: T
    message: string
    success: boolean
}

interface dataRequestParams {
    url: string
    data?: FormData | string | object | (() => any)
    loading?: () => void // 等待開使
    loadingActionName?: string // 載入的 action name
}

interface formRequestParams extends dataRequestParams, SweetAlertOptions {
    check_title?: string
    check_type?: 'question' | 'warning' | 'error' | 'success' | 'info'
    check_text?: string
    imgUrl?: string
    method?: 'put' | 'post' | 'delete' | 'get'
}

// instance
const instance = axios.create({
    withCredentials: true,
    baseURL: 'http://localhost:4000/'
})

const showErrorDialog = async (error): Promise<void> => {
    const data = error?.response?.data || error?.data
    await MySwal.AlertMessage({
        icon: 'error',
        title: data?.message
    })
}

// 請求前攔截器
instance.interceptors.request.use(
    async (config) => {
        return config
    },
    async (error) => {
        // this.loading('onLoading', false)
        void showErrorDialog(error)
        return await Promise.reject()
    }
)

// 回傳攔截器
instance.interceptors.response.use(
    (response): any => {
        // this.loading('onLoading', false)
        return response?.data
    },
    async (error) => {
        // this.loading('onLoading', false)
        void showErrorDialog(error)
        return await Promise.reject()
    }
)

/**
 * 當繼承 此類的時候
 * 請給此類 一個 redux store and loadingAction.
 * 如此才會照常 執行
 *
 * @route 路由 {}
 * @type api ctl type {}
 * @checkTitle
 * @loadingAction
 * @store
*/
class apiAbstract {
    // api type
    public type: Record<string, any>

    protected checkTitle: Record<string, any>

    protected loadingAction = allAction.loading

    protected action = allAction

    protected store = store

    private loading (loadingActionName: string | undefined, state: boolean): void {
        let actionName = loadingActionName as unknown as string

        // 如果沒有 loadingActionName，就轉成 'loading'
        if (!actionName || !this.loadingAction?.[actionName]) {
            actionName = 'onLoading'
        }
        const LA = (
            this.loadingAction?.[actionName] ||
            (() => {})
        )(state);

        (
            this.store?.dispatch ||
            (() => {})
        )(LA)
    }

    // 請求 這是有 確認框的
    /**
     * @T 是 回傳類型
     * @description 這個方法會驗證兩次form 表單 因為 check alert window 裡面也有可能放表單
    */
    private async ConfirmRequest <T extends any>(v: formRequestParams): Promise<T> {
        return await MySwal.checkMessage<responseParams<T>>({
            title: v?.check_title,
            icon: v?.check_type || 'question',
            text: v?.check_text,
            imageUrl: v?.imgUrl,
            showLoaderOnConfirm: true,
            ...v,
            preConfirm: async () => {
                return await instance?.[v?.method]<T>(v?.url, v?.data)
            }
        }).then((response): T => {
            void MySwal.AlertMessage({
                icon: 'success',
                title: response?.value?.message
            })
            return response?.value?.data as unknown as T
        })
    }

    public async POST<T extends any>(v: formRequestParams): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'post' })
    }

    public async PUT<T extends any>(v: formRequestParams): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'put' })
    }

    public async DELETE<T extends any>(v: formRequestParams): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'delete' })
    }

    // 請求 這是沒有確認框的
    /**
     * @T 是 回傳類型
    */
    public async GET <T extends any>(v: dataRequestParams): Promise<T> {
        return await instance.get<T>(
            v?.url, {
                params: v?.data
            }).then((v) => v?.data)
    }
}

export default apiAbstract
