import { Avatar, List, Tabs } from 'antd'
import Search from 'antd/lib/input/Search'
import React, { useEffect, useRef, useState } from 'react'
import api from '../../api/api'
import Btn from '../../component/Btn'
import useReduceing from '../../Hook/useReducing'
import { CompanyType } from '../../type'
import companyIcon from '../../asserts/company.png'
const HomePage = (): JSX.Element => {
    const { company, loading } = useReduceing()
    const [sign, setSign] = useState<CompanyType[]>([])
    const companyCodeRef = useRef('')
    const onSearch = async (): Promise<void> => {
        await api.getCompanyInfo({ companyCode: companyCodeRef.current })
    }
    const onThreeTab = async (): Promise<void> => {
        const res = await api.getWaitReply()
        if (res.status && res.data?.length > 0) {
            const waitReply = res.data
            for (let i = 0; i < waitReply?.length; i++) {
                const companys = await api.getCompanyInfo({ companyId: waitReply[i].CompanyId })
                setSign((prev) => ([...prev, ...companys.data]))
            }
        }
    }
    const onJoin = async (companyCode: CompanyType['CompanyCode']): Promise<void> => {
        await api.createWaitReply(companyCode)
    }
    useEffect(() => {
        api.getSelfData()
        api.getCompanyInfo({ companyCode: '' })
    }, [])
    return (
        <div className={styles.HomeBlock}>
            <Tabs onChange={(e) => {
                setSign([])
                if (e === '3') {
                    onThreeTab()
                }
                onSearch()
            }}>
                <Tabs.TabPane tab='加入公司' key={1}>
                    <Search
                        onChange={(e) => {
                            companyCodeRef.current = e.target.value
                        }}
                        onSearch={onSearch}
                        placeholder="input search text"
                        enterButton="Search"
                        size="large"
                        loading={loading.onFetchCompany}
                    />
                    <List
                        loading={loading.onFetchCompany}
                        renderItem={(item: CompanyType) => {
                            return (
                                <List.Item style={{ flexWrap: 'wrap', position: 'relative' }}>
                                    <List.Item.Meta
                                        style={{ minWidth: '200px', marginBottom: '30px' }}
                                        avatar={<Avatar src={companyIcon} />}
                                        title={<>{item.CompanyName}</>}
                                        description={`公司地址: ${item.CompanyLocation}`}
                                    />
                                    <div style={{ position: 'absolute', right: '0px', bottom: '5px' }}>
                                        <Btn.Confirm onClick={() => {
                                            onJoin(item.CompanyCode)
                                        }} />
                                    </div>

                                </List.Item>
                            )
                        }}
                        dataSource={Object.keys(company?.info)?.length > 0 ? [company.info] : []}
                    />
                </Tabs.TabPane>
                <Tabs.TabPane tab='創建公司' key={2}>
                    <></>
                </Tabs.TabPane>
                <Tabs.TabPane tab='申請中' key={3}>
                    <List
                        loading={loading.onFetchCompany || loading.onFetchWaitReply}
                        renderItem={(item: CompanyType) => {
                            return (
                                <List.Item style={{ flexWrap: 'wrap', position: 'relative' }}>
                                    <List.Item.Meta
                                        style={{ minWidth: '200px', marginBottom: '30px' }}
                                        avatar={<Avatar src={companyIcon} />}
                                        title={<>{item.CompanyName}</>}
                                        description={<>
                                            公司地址: {item.CompanyLocation}<br/>
                                            公司碼: {item.CompanyCode}<br/>
                                        </>
                                        }
                                    />
                                    <div style={{ position: 'absolute', right: '0px', bottom: '5px' }}>
                                        <Btn.Cancel onClick={() => {
                                        }} />
                                    </div>

                                </List.Item>
                            )
                        }}
                        dataSource={sign}
                    />
                </Tabs.TabPane>
            </Tabs>
        </div>
    )
}
export default HomePage
