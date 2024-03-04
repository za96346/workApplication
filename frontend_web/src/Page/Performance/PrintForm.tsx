/* eslint-disable no-tabs */
import { useAppSelector } from 'hook/redux'
import React, { useEffect, useRef, ReactNode, useState } from 'react'
import performanceTypes from 'types/performance'

const breakPoint = {
    signBlock: [12, 15],
    trackListBlock: [15, 17.5],
    feedBackBlock: [17.5, 22.5],
    workBlock: [22.5, 27.5],
    memberFeedBackBlock: [27.5, '']
}

const Form = ({ value }: { value: performanceTypes.reducerType['all'][0] }): JSX.Element => {
    const [breakState, setBreakState] = useState<Record<string, 'always' | 'avoid'>>({
        signBlock: 'avoid',
        trackListBlock: 'avoid',
        feedBackBlock: 'avoid',
        workBlock: 'avoid',
        memberFeedBackBlock: 'avoid'
    })
    const isPageBreakBeforeRef = useRef(false)
    const selfEvaluationRef = useRef<HTMLDivElement>()
    const monthPerformanceRef = useRef<HTMLDivElement>()

    const calBreak = (num1: any, num2: any, key): 'always' | 'avoid' => {
        const selfEvaluationHeight = selfEvaluationRef.current?.offsetHeight ?? 0
        const monthPerformanceHeight = monthPerformanceRef.current?.offsetHeight ?? 0

        const heightTotal = (selfEvaluationHeight + monthPerformanceHeight) / 38

        console.log('heightTotal => ', heightTotal)

        const result = heightTotal >= num1 && heightTotal < num2
            ? 'always'
            : 'avoid'

        if (!isPageBreakBeforeRef.current) {
            setBreakState((prev) => ({ ...prev, [key]: 'always' }))
        }

        isPageBreakBeforeRef.current = true

        return result
    }

    const handleGoal = (goal: string): ReactNode => {
        return goal.split('\n').map((item) => (
            item.length === 0 || item === ''
                ? <></>
                : (
                    <>
                        &nbsp;&nbsp;&nbsp; {item}&nbsp;&nbsp;&nbsp;&nbsp;口已完成  口目標持續中<br/>
                    </>
                )
        ))
    }

    useEffect(() => {
        setTimeout(() => {
            calBreak(
                breakPoint.signBlock[0],
                breakPoint.signBlock[1],
                'signBlock'
            )
            calBreak(
                breakPoint.trackListBlock[0],
                breakPoint.trackListBlock[1],
                'trackListBlock'
            )
            calBreak(
                breakPoint.feedBackBlock[0],
                breakPoint.feedBackBlock[1],
                'feedBackBlock'
            )
            calBreak(
                breakPoint.workBlock[0],
                breakPoint.workBlock[1],
                'workBlock'
            )
            calBreak(
                breakPoint.memberFeedBackBlock[0],
                breakPoint.memberFeedBackBlock[1],
                'memberFeedBackBlock'
            )
        }, 200)
    }, [])

    return (
        <div translate='no' className="print_page">
            <p style={{ fontSize: '8px' }}>臺中市私立鎮瀾兒童家園</p>
            <div style={{ fontSize: '20px', justifyContent: 'center', border: 'none', marginBottom: '0.5cm' }}>工作督導月紀錄</div>
            <div style={{ height: '1cm' }}>
                <div className='d-flex justify-content-center align-items-center' style={{ width: '4.5cm' }} >日期</div>
                <div className='d-flex justify-content-center align-items-center' style={{ width: '4.5cm' }}>
                    {value.Year}-{value.Month < 10 ? `0${value.Month}` : value.Month}
                </div>
                <div className='d-flex justify-content-center align-items-center' style={{ width: '4.5cm' }}>姓名</div>
                <div className='d-flex justify-content-center align-items-center' style={{ width: '4.5cm' }}>{value?.UserName || ''}</div>
            </div>
            <div
                ref={selfEvaluationRef}
                style={{ height: 'fit-content' }}
            >
                <div className={'print_page_column_one'}>工作自評</div>
                <div className={'print_page_column_two'}>
                    1、	紀錄繳交 □依規定完成  □遲交  □未完成<br/>
                    2、	方案執行 □依規定完成  □遲交  □未完成<br/>
                    3、	個人目標<br/>
                    {handleGoal(value.Goal)}
                </div>
            </div>
            <div
                ref={monthPerformanceRef}
                style={{ minHeight: '6cm', height: 'fit-content' }}>
                <div className={'print_page_column_one'}>當月績效</div>
                <div className={'print_page_column_two'}>
                    專業績效:{value?.Professional || 0}&nbsp;    行政效率:{value?.Efficiency || 0}&nbsp;    工作態度:{value?.Attitude || 0}<br/>
                    出勤狀況&nbsp;&nbsp;&nbsp;遲到:{value?.BeLate || 0}&nbsp;    未依規定請假:{value?.DayOffNotOnRule || 0}<br/>
                    績效描述:
                    <br/>
                    {
                        value.Directions?.split('\n').map((item) => {
                            return (<>{item}<br/></>)
                        })
                    }
                </div>
            </div>
            <div style={{
                height: '2.5cm',
                marginTop: breakState.memberFeedBackBlock === 'always' ? '1.5cm' : '0px',
                pageBreakBefore: breakState.memberFeedBackBlock
            }}>
                <div className={'print_page_column_one'}>組員回饋</div>
                <div className={'print_page_column_two'}></div>
            </div>
            <div style={{
                height: '5cm',
                marginTop: breakState.workBlock === 'always' ? '1.5cm' : '0px',
                pageBreakBefore: breakState.workBlock
            }}>
                <div className={'print_page_column_one'}>工作督導</div>
                <div className={'print_page_column_two'}>
                    一、組員議題（可複選） <br/>
                    &nbsp;&nbsp;&nbsp; □理論運用 □專業方法 □專業倫理 □專業關係 □督導關係 □同儕關係<br/>
                    &nbsp;&nbsp;&nbsp; □機構運作 □價值衝突 □行政業務 □個人自我突破與成長<br/>
                    &nbsp;&nbsp;&nbsp; □其他（請說明）________________________________  <br/>

                    二、督導之處置（可複選） <br/>
                    &nbsp;&nbsp;&nbsp; □專業概念解說與討論 □情緒支持與同理 □與機構主管溝通 <br/>
                    &nbsp;&nbsp;&nbsp; □協助組員進行反思 □提供閱讀資源  □提供個人實習經驗<br/>
                    &nbsp;&nbsp;&nbsp; □向園方反應 □要求組員達成園方基本要求 <br/>
                    &nbsp;&nbsp;&nbsp; □其他（請說明）________________________________
                </div>
            </div>
            <div style={{
                height: '2.25cm',
                marginTop: breakState.feedBackBlock === 'always' ? '1.5cm' : '0px',
                pageBreakBefore: breakState.feedBackBlock
            }}>
                <div className={'print_page_column_one'}>督導回饋</div>
                <div className={'print_page_column_two'}></div>
            </div>
            <div style={{
                height: '2.25cm',
                marginTop: breakState.trackListBlock === 'always' ? '1.5cm' : '0px',
                pageBreakBefore: breakState.trackListBlock
            }}>
                <div className={'print_page_column_one'}>追蹤事項</div>
                <div className={'print_page_column_two'}></div>
            </div>
            <div
                style={{
                    height: '1cm',
                    marginTop: breakState.signBlock === 'always' ? '1.5cm' : '0px',
                    pageBreakBefore: breakState.signBlock
                }}>
                <div className='d-flex align-items-center' style={{ width: '9cm' }}>記錄者簽章:</div>
                <div className='d-flex align-items-center' style={{ width: '9cm' }}>督導簽章:</div>
            </div>
            <div className='d-flex justify-content-end' style={{ fontSize: '8px' }}>
                1080101第三次修定
            </div>
        </div>
    )
}

const PrintForm = (): JSX.Element => {
    const performance = useAppSelector((v) => v?.performance?.all)
    useEffect(() => {
        setTimeout(() => {
            print()
        }, 1000)
    }, [])
    return (
        <>
            {
                performance.map((value) => (
                    <Form key={value?.PerformanceId} value={value} />
                ))
            }
        </>
    )
}
export default PrintForm
