import { session } from 'hook/useSession'
import systemTypes from 'types/system'

const Session = session<Partial<systemTypes.auth['permission']>>({})
export default Session
