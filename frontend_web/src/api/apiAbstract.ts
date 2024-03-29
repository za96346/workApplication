/* eslint-disable prefer-promise-reject-errors */
import axios, { AxiosRequestConfig, AxiosResponse } from 'axios'
import { type SweetAlertOptions } from 'sweetalert2'
// import { reducer } from '@vteam_components/redux'
import { store, allAction } from 'reducer/store'

// 這個 global 一定要引入 不然此api build 時候會讀不到 actionT and storeT
// import '@vteam_components/redux/global'
import { MySwal } from '../func/Swal'
import { v4 } from 'uuid'

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
    checkTitle?: string
    check_type?: 'question' | 'warning' | 'error' | 'success' | 'info'
    checkText?: string
    imgUrl?: string
    method?: 'put' | 'post' | 'delete' | 'get'
}

interface formValueType {
    formName: string
    validCheck: boolean
    InvalidAction?: () => void
}

type axiosRequestConfig = AxiosRequestConfig & { id: string }
type axiosResponse = AxiosResponse<any, any> & { config: { id: string } }

// instance
const instance = axios.create({
    withCredentials: true,
    baseURL: process.env.REACT_APP_API
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
    async (config: axiosRequestConfig) => {
        config.id = v4()
        store.dispatch(allAction.loading.onLoading({ [config.id]: true }))
        return config
    },
    async (error) => {
        store.dispatch(allAction.loading.onLoading({ [error.config.id]: false }))
        void showErrorDialog(error)
        return await Promise.reject()
    }
)

// 回傳攔截器
instance.interceptors.response.use(
    (response: axiosResponse): any => {
        store.dispatch(allAction.loading.onLoading({ [response.config.id]: false }))
        return response?.data
    },
    async (error) => {
        store.dispatch(allAction.loading.onLoading({ [error.config.id]: false }))
        void showErrorDialog(error)

        if (error.response.status === 511) {
            window.location.href = '/entry/login'
        }

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

    protected checkTitle = {
        confirmUpdate: '確定修改？',
        confirmAdd: '確定新增？',
        confirmDelete: '確認刪除？'
    }

    protected loadingAction = allAction.loading

    protected action = allAction

    protected store = store

    // 生成 form
    /**
     * @param formValue = [
     *      {
     *          formName: string, // dont need a hashTag
     *          validCheck: boolean,
     *          InvalidAction: () => {}
     *      }
     * ]
     * @return FormData
    */
    public makeFormData (formValue: formValueType[]): [Record<string, any>, boolean] {
        const formDataObject = {}
        let isValid = true
        formValue.forEach((item) => {
            const form: HTMLFormElement | null = document.querySelector(`#${item?.formName}`)
            if (form) {
                // 如果 需要檢查 以及 檢查沒過
                if (item.validCheck && !form?.checkValidity()) {
                    // 先跑 原生跳出
                    // 如果 發現 找步道 就換 tabs emit
                    form?.reportValidity()

                    if (item?.InvalidAction) item?.InvalidAction()
                    isValid = false
                    throw new Error('表單驗證失敗')
                }
                const a = new FormData(form)
                // eslint-disable-next-line no-restricted-syntax
                for (const pair of a.entries()) {
                    formDataObject[pair[0]] = pair[1]
                }
            }
        })
        return [formDataObject, isValid]
    }

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
            title: v?.checkTitle,
            icon: v?.check_type || 'question',
            text: v?.checkText,
            imageUrl: v?.imgUrl,
            showLoaderOnConfirm: true,
            ...v,
            preConfirm: async () => {
                return await instance?.[v?.method]<T>(
                    v?.url,
                    v?.method === 'delete'
                        ? {
                            data: v?.data
                        }
                        : v?.data
                )
            }
        }).then((response): T => {
            void MySwal.AlertMessage({
                icon: 'success',
                title: response?.value?.message
            })
            return response?.value?.data as unknown as T
        })
    }

    protected async POST<T extends any>(v: formRequestParams): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'post' })
    }

    protected async PUT<T extends any>(v: formRequestParams): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'put' })
    }

    protected async DELETE<T extends any>(v: formRequestParams): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'delete' })
    }

    // 請求 這是沒有確認框的
    /**
     * @T 是 回傳類型
    */
    protected async GET <T extends any>(v: dataRequestParams): Promise<T> {
        return await instance.get<T>(
            v?.url,
            {
                params: v?.data
            }
        ).then((v) => v?.data)
    }
}

export default apiAbstract
