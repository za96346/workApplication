import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import CompanyForm from './component/form/Company'
import PersonalForm from './component/form/Personal'
import statics from '../../statics'
import api from 'api/api'
import { Modal } from 'antd'
import PersonalDescibe from './component/describe/Personal'
import CompanyDescibe from './component/describe/Company'
// import Btn from 'Share/Btn'

const statusInit = {
    changePwdBtn: false,
    onChangePwd: false
}

const Index = (): JSX.Element => {
    let { types } = useParams()
    types = types.replace('z', '')
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [status, setStatus] = useState(statusInit)

    useEffect(() => {
        if (types === statics.personalSetting) {
            void api.getSelfData()
        }
    }, [types])
    return (
        <>
            <Modal
                forceRender
                closeIcon={<></>}
                footer={null}
                open={status.changePwdBtn}
            >
                <CompanyForm />
                <PersonalForm />
            </Modal>
            <div className={window.styles.settingBlock}>
                {/* 個人資料設定 */}
                {
                    types === statics.personalSetting && (
                        <>
                            <PersonalDescibe />
                        </>
                    )
                }
                {/* 公司資料設定 */}
                {
                    types === statics.companySetting && (
                        <>
                            <CompanyDescibe />
                        </>
                    )
                }
                {/* <Btn.Edit onClick={() => {}}/> */}
            </div>
        </>
    )
}
export default Index
