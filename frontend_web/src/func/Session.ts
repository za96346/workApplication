import { v4 } from 'uuid'
import listenerAbstract from './ListenerAbs'

interface sessionType<T> {
    /* 設定值 */
    set: (v: T | ((prev: T) => T)) => void

    /* 獲取值 */
    get: (callback: (target: T, property: keyof T) => void) => T

    /* 往後尋找 */
    forward: (step: number) => void

    /* 往前尋找 */
    backward: (step: number) => void

    reset: (resetValue?: T) => void
}

/**
 * @description 會話 可以存儲資料，並且可以搭配provider 使用
 * @constructor {
 *      initValue = 預設值
 * }
*/
class Session<T extends {}> extends listenerAbstract implements sessionType<T> {
    #data: T // 此物件直

    #previousDataQuene: any[] = [] // 歷史資料的隊列

    #previousDataQuenePointer: number = 0 // 當前的指標 ( array idx )

    #allowStorage: boolean = false

    #initValue: T

    public readonly id: string

    constructor (
        // eslint-disable-next-line default-param-last
        initValue = {},
        {
            id = v4(),
            allowStorage = false
        }
    ) {
        super()
        this.id = id

        if (allowStorage) {
            const storageData = JSON.parse(sessionStorage.getItem(this.id))
            this.#data = storageData?.data || initValue
            this.#previousDataQuene = storageData?.previousData || []

            sessionStorage.setItem(
                this.id,
                JSON.stringify({
                    data: this.#data,
                    previousData: this.#previousDataQuene
                })
            )
        } else {
            this.#data = (initValue || {}) as T
            this.#previousDataQuene.push(this.#data)
        }

        this.#initValue = (initValue || {}) as T
        this.#allowStorage = allowStorage || false

        this.#previousDataQuenePointer = (this.#previousDataQuene?.length || 1) - 1
    }

    #setData (v): void {
        this.#data = typeof v === 'function'
            ? v(this.#data)
            : v

        this.emit(
            this.id,
            v
        )
    }

    /* 往前尋找 */
    public backward (step: number): void {
        const position = this.#previousDataQuenePointer - step

        this.#previousDataQuenePointer = position < 0 ? 0 : position

        this.#setData(this.#previousDataQuene[this.#previousDataQuenePointer])
    }

    /* 往後尋找 */
    public forward (step: number): void {
        const position = this.#previousDataQuenePointer + step

        this.#previousDataQuenePointer = position < 0 ? 0 : position

        this.#setData(this.#previousDataQuene[this.#previousDataQuenePointer])
    }

    /* 設定值 */
    public set (v: T | ((prev: T) => T)): void {
        this.#setData(v)
        this.#previousDataQuene.splice(
            ++this.#previousDataQuenePointer,
            0,
            JSON.parse(JSON.stringify(this.#data))
        )

        // 如果允許 儲存在本地的時候 就把資料 和 歷史資料存到 sessionStorage
        if (this.#allowStorage) {
            sessionStorage.setItem(
                this.id,
                JSON.stringify({
                    data: this.#data,
                    previousData: this.#previousDataQuene
                })
            )
        }
    }

    /* 獲取值 */
    public get (proxyListenerCallback = (target, property) => {}): T {
        return new Proxy(this.#data, {
            get: (target, property, receiver) => {
                proxyListenerCallback(target, property)
                return target[property]
            }
        })
    }

    /* 重製狀態 */
    public reset (resetValue?: T): void {
        sessionStorage.removeItem(this.id)
        this.#data = resetValue || this.#initValue
        this.#previousDataQuene = []
        this.#previousDataQuenePointer = 0
    }

    public override on (type, callback): void {
        super.on(this, type, callback)
    }
}
export default Session
