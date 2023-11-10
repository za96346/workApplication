/* eslint-disable prefer-promise-reject-errors */
import React from 'react'
import axios, { type AxiosInstance } from 'axios'
import { type SweetAlertOptions } from 'sweetalert2'
// import { reducer } from '@vteam_components/redux'
import { store, allAction } from 'reducer/store'

// 這個 global 一定要引入 不然此api build 時候會讀不到 actionT and storeT
// import '@vteam_components/redux/global'
import { MySwal } from '../func/Swal'

interface responseParams<T> {
    data: T
    message: string
    success: boolean
}

interface dataRequestParams<T> {
    url: string
    data?: FormData | string | object | (() => any)
    cb?: (response: T) => void // 拿到 response 成功
    cp?: () => void // 不管 成功失敗 都執行的 區塊
    failModal?: (response: responseParams<T>) => void // 失敗 modal callback(覆蓋掉 原本的)
    reject?: (reason?: any) => void
    loading?: () => void // 等待開使
    loadingActionName?: string // 載入的 action name
}

interface formRequestParams<T> extends dataRequestParams<T>, SweetAlertOptions {
    check_title?: string
    check_type?: 'question' | 'warning' | 'error' | 'success' | 'info'
    check_text?: string
    imgUrl?: string
    method: 'put' | 'post' | 'delete' | 'get'
    successModal?: (response: responseParams<T>) => void // 成功 modal callback (覆蓋掉 原本的)
}

interface formValueType {
    formName: string
    validCheck: boolean
    InvalidAction: () => void
}

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
abstract class apiAbstract {
    // api type
    public type: Record<string, any>

    protected checkTitle: Record<string, any>

    protected loadingAction = allAction.loading

    protected action = allAction

    protected store = store

    protected axios: AxiosInstance

    /**
     * @description 登入路由 當 請求 session 過期時 會導到此路由
    */
    protected loginRoute: string = '/'

    constructor () {
        const token = document.head.querySelector('meta[name="csrf-token"]')

        this.axios = axios.create({
            headers: {
                common: {
                    'X-CSRF-TOKEN': (token as any)?.content || ''
                }
            },
            timeout: 30000
        })

        // axiosRetry(this.axios, {
        //     retries: 4,
        //     retryCondition: () => true,
        //     retryDelay: () => 1000,
        // })

        // 請求前攔截器
        this.axios.interceptors.request.use(
            async (config) => {
                const originParams: dataRequestParams<any> = { ...(config?.data || {}) }
                config.data = config.data.data

                return { ...config, originParams }
            },
            async (error) => {
                void this.showErrorDialog(error)
                return await Promise.reject(error)
            }
        )

        // 回傳攔截器
        this.axios.interceptors.response.use(
            (response): any => {
                const data: responseParams<any> = response?.data
                if (!data?.success) {
                    void this.showErrorDialog(response)
                    return Promise.reject(data || '')
                }
                return data
            },
            async (error) => {
                void this.showErrorDialog(error)
                return await Promise.reject(error?.response?.data)
            }
        )
    }

    private async showErrorDialog (error): Promise<void> {
        const { originParams } = error.config
        const data = error?.response?.data || error?.data
        if (data?.data?.redirect) {
            window.location.href = this.loginRoute
            return
        }
        if (!originParams?.failModal) {
            await MySwal.AlertMessage({
                icon: 'error',
                title: data?.message
            })
        } else {
            originParams.failModal(data)
        }
    }

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
    public makeFormData (formValue: formValueType[]): [FormData, boolean] {
        const formData = new FormData()
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
                    formData.append(pair[0], pair[1])
                }
            }
        })
        return [formData, isValid]
    }

    private loading (loadingActionName: string | undefined, state: boolean): void {
        const La = this.loadingAction as Record<string, any>
        let actionName = loadingActionName as unknown as string

        // 如果沒有 loadingActionName，就轉成 'loading'
        if (!loadingActionName || !La?.[loadingActionName]) {
            actionName = 'onLoading'
        }
        const LA = (
            La?.[actionName] ||
            (() => {})
        )(state);

        (
            this.store?.dispatch ||
            (() => {})
        )(LA)
    }

    // 驗證資料
    private validateData<T>(v: dataRequestParams<T> | formRequestParams<T>): void {
        try {
            v.data = typeof v.data === 'function'
                ? v.data()
                : v.data
        } catch (e) {
            throw new Error(e)
        }
    }

    // 請求 這是沒有確認框的
    /**
     * @T 是 回傳類型
    */
    protected async GET <T extends Record<any, any>>(v: dataRequestParams<T>): Promise<T> {
        return await this.axios.get<T>(v?.url)
            .then((response) => response?.data)
    }

    // 請求 這是有 確認框的
    /**
     * @T 是 回傳類型
     * @description 這個方法會驗證兩次form 表單 因為 check alert window 裡面也有可能放表單
    */
    private async ConfirmRequest <T extends Record<any, any>>(v: formRequestParams<T>): Promise<T> {
        return await MySwal.checkMessage<responseParams<T>>({
            title: v?.check_title,
            icon: v?.check_type || 'question',
            text: v?.check_text,
            imageUrl: v?.imgUrl,
            showLoaderOnConfirm: true,
            ...v,
            // eslint-disable-next-line consistent-return
            preConfirm: async () => {
                try {
                    this.validateData<T>(v)
                    return await this.axios?.[v?.method]<T>(v?.url, v)
                } catch {
                    MySwal.MySwal.showValidationMessage('驗證失敗')
                }
            }
        }).then((response): T => {
            if (v?.successModal) {
                v?.successModal(response.value as NonNullable<typeof response.value>)
            } else {
                void MySwal.AlertMessage({
                    icon: 'success',
                    title: response?.value?.message
                })
            }
            return response?.value?.data as unknown as T
        })
    }

    protected async POST<T extends Record<any, any>>(v: formRequestParams<T >): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'post' })
    }

    protected async PUT<T extends Record<any, any>>(v: formRequestParams<T >): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'put' })
    }

    protected async DELETE<T extends Record<any, any>>(v: formRequestParams<T >): Promise<T> {
        return await this.ConfirmRequest({ ...v, method: 'delete' })
    }
    
}
export default apiAbstract
