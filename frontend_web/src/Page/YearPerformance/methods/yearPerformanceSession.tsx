import { performanceParams } from 'api/Index'
import { session } from 'hook/useSession'
import { FuncCodeEnum } from 'types/system'

interface yearPerformanceSessionType {
    currentParams?: performanceParams.getYear
}

const YearPerformanceSession = session<yearPerformanceSessionType>({
    currentParams: {}
}, {
    id: FuncCodeEnum.yearPerformance
})

export default YearPerformanceSession
