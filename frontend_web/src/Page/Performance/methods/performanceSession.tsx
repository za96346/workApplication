import { performanceParams } from 'api/Index'
import { session } from 'hook/useSession'
import { FuncCodeEnum } from 'types/system'

interface performanceSessionType {
    currentParams?: performanceParams.get
}

const PerformanceSession = session<performanceSessionType>({
    currentParams: {}
}, {
    id: FuncCodeEnum.performance
})

export default PerformanceSession
