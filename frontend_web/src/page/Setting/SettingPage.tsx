import React, { useEffect } from 'react'
import { useParams } from 'react-router-dom'
import CompanyForm from './CompanyForm'
import PersonalForm from './PersonalForm'
import statics from '../../statics'
import api from '../../api/api'

const SettingPage = (): JSX.Element => {
    const { types } = useParams()

    useEffect(() => {
        if (types === statics.personalSetting) {
            void api.getSelfData()
        }
    }, [types])
    return (
        <>
            <div className={styles.settingBlock}>
                {/* 個人資料設定 */}
                {
                    types === statics.personalSetting && (
                        <>
                            <div className={styles.settingBody}>
                                <div>編輯個人資料</div>
                                <div>
                                    <PersonalForm />
                                </div>
                            </div>

                        </>
                    )
                }
                {/* 公司資料設定 */}
                {
                    types === statics.companySetting && (
                        <>
                            <div className={styles.settingBody}>
                                <div>公司資料</div>
                                <div>
                                    <CompanyForm />
                                </div>
                            </div>
                        </>
                    )
                }
            </div>
        </>
    )
}
export default SettingPage
