/* eslint-disable no-underscore-dangle */

type eventType = string

abstract class listenerAbstract {
    private event: Record<eventType, any[]>

    constructor () {
        this.event = {}
    }

    // 發送事件
    protected emit (type: eventType, message): void {
        if (this.event[type]) {
            this.event[type].forEach((item) => {
                item.callback(message)
            })
        }
    }

    // 監聽
    protected on (objs, type: eventType, callback): void {
        this.event[type] = this.event[type] || []
        this.event[type].push({
            objs,
            callback
        })
    }

    // 關閉
    protected off (obj, type: eventType): void {
        if (this.event[type]) {
            this.event[type] = this.event[type].filter((item) => item.obj !== obj)
        }
    }
}
export default listenerAbstract
