import { BanchType } from "../../type"
import { action, companyType } from "../types"

class companyAction {
    constructor () {
        this.setBanch = this.setBanch.bind(this)
    }

    setBanch (banch: BanchType[]): action {
        return {
            type: companyType.SET_BANCH,
            payload: {
                banch
            }
        }
    }
}
export default new companyAction()
