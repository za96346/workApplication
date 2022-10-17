/* eslint-disable class-methods-use-this */
import moment from 'moment'

interface DateHandleInterface {
    now: Date
    nowDay: any
    nowMonth: any
    nowYear: any
    dateFormatOne: string
    dateFormatTwo: string
    formatDate: Function
    fullFormat: string
    getMonthDays: Function
    getWeekStartDate: Function
    getWeekEndDate: Function
    getMonthStartDate: Function
    getMonthEndDate: Function
    transferToUtc: Function
    renderFirst: Function
}

type DateOrString = string | Date

class DateHandle implements DateHandleInterface {
    now = new Date() // 当前日期

    nowDayOfWeek = this.now.getDay() - 1 // 今天本周的第几天

    nowDay = this.now.getDate() // 当前日

    nowMonth = this.now.getMonth() // 当前月

    nowYear = this.now.getFullYear() // 当前年

    dateFormatOne = 'YYYY/MM/DD'

    dateFormatTwo = 'YYYY-MM-DD'

    timeFormat = ' HH:mm:ss'

    fullFormat = 'YYYY/MM/DD HH:mm:ss'

    constructor () {
        this.nowYear += this.nowYear < 2000 ? 1900 : 0
        this.formatDate = this.formatDate.bind(this)
        this.getMonthDays = this.getMonthDays.bind(this)
        this.getWeekEndDate = this.getWeekEndDate.bind(this)
        this.getWeekStartDate = this.getWeekStartDate.bind(this)
        this.getMonthStartDate = this.getMonthStartDate.bind(this)
        this.getMonthEndDate = this.getMonthEndDate.bind(this)
        this.formatNow = this.formatNow.bind(this)
    }

    // 格式化日期：yyyy-MM-dd
    formatDate (date: any): any {
        // console.log(date)
        const myyear = date.getFullYear()
        // eslint-disable-next-line @typescript-eslint/restrict-plus-operands
        let mymonth = date.getMonth() + 1
        let myweekday = date.getDate()
        if (mymonth < 10) {
            mymonth = `0${mymonth}`
        }
        if (myweekday < 10) {
            myweekday = `0${myweekday}`
        }
        return `${myyear}-${mymonth}-${myweekday}`
    }

    // 获得某月的天数
    getMonthDays (myMonth: any): any {
        const monthStartDate: any = new Date(this.nowYear, myMonth, 1)
        // eslint-disable-next-line @typescript-eslint/restrict-plus-operands
        const monthEndDate: any = new Date(this.nowYear, myMonth + 1, 1)
        const days = (monthEndDate - monthStartDate) / (1000 * 60 * 60 * 24)
        return days
    }

    // 获得本周的开始日期
    getWeekStartDate (): any {
        const weekStartDate = new Date(this.nowYear, this.nowMonth, this.nowDay - this.nowDayOfWeek)
        return this.formatDate(weekStartDate)
    }

    // 获得本周的结束日期
    getWeekEndDate (): any {
        const weekEndDate = new Date(
            this.nowYear,
            this.nowMonth,
            this.nowDay + (6 - this.nowDayOfWeek)
        )
        return this.formatDate(weekEndDate)
    }

    // 获得本月的开始日期
    getMonthStartDate (): any {
        const monthStartDate = new Date(this.nowYear, this.nowMonth, 1)
        return this.formatDate(monthStartDate)
    }

    // 获得本月的结束日期
    getMonthEndDate (): any {
        const monthEndDate = new Date(this.nowYear, this.nowMonth, this.getMonthDays(this.nowMonth))
        return this.formatDate(monthEndDate)
    }

    transferToUtc (date: DateOrString, time: DateOrString): any {
    // console.log(date, time, `utc => ${new Date(`${date} ${time}`).toISOString()}`);
        const res = `${date} ${time}`.replace(/-/g, '/').replace('T', ' ')
        return new Date(res).toISOString()
    }

    dateFormatToTime (date: Date): string {
        let hours: number | string = date.getHours()
        let min: number | string = date.getMinutes()
        if (hours <= 9) {
            hours = `0${hours}`
        } else {
            hours = hours.toString()
        }
        if (min <= 9) {
            min = `0${min}`
        } else {
            min = min.toString()
        }
        return `${hours}:${min}:00`
    }

    transferUtcFormat (date: any, offset?: number): any {
        return moment(date)
            .utcOffset(offset || 8)
            .format(this.fullFormat)
    }

    renderFirst (): any {
        console.log('it me')
        const start = this.transferToUtc(this.formatDate(this.now), '00:00:00')
        const end = this.transferToUtc(this.formatDate(this.now), '23:59:59')
        // console.log(this.now, start, end);
        return [start, end]
    }

    formatNow (): any {
    // console.log(this.formatDate(this.now));
        return this.formatDate(this.now)
    }

    diffYears (startDate: DateOrString, endDate: DateOrString): any {
        return moment(endDate).diff(moment(startDate), 'years')
    }
}

export default new DateHandle()
