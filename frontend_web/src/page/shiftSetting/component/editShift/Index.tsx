import React, { useEffect } from "react";
import { Form, List, listType, SelectUI } from "@vteam_components/cloud";
import { Avatar, Divider } from "antd";
import api from "api/api";
import useReduceing from "Hook/useReducing";
import { UserType } from "type";
import statics from 'statics';

const Index = (): JSX.Element => {
    const { state, company } = useReduceing()
    useEffect(() => {
        api.getUserAll({
            workState: 'on',
            banch: state.banchId 
        });
    }, [state?.banchId])
    return (
        <>
            <div className="alert alert-warning">
                編輯後為 立即生效。
            </div>
            <Divider>基本設定</Divider>
            <Form onChange={(v) => { console.log(v) }}>
                <Form.Select
                    className="col-md-6"
                    name="editMode"
                    datas={statics.shiftSettingObj}
                    title="編輯模式"
                    defaultValue={'coEdit'}
                    onChange={() => {}}
                />
                <Form.Select
                    className="col-md-6"
                    name="editScope"
                    datas={statics.scope}
                    title="編輯範圍"
                    defaultValue={'self'}
                    onChange={() => {}}
                />

                {/* 排序 */}
                <SelectUI onChange={(v) => { console.log(v) }}>
                    <SelectUI.Group
                        groupName='group'
                        required={false}
                        maxChecked={company.employee?.length || 0}
                    >
                        <Form.Status>
                            {
                                ({ fieldValue }) => (fieldValue?.editMode === 'sortEdit' || fieldValue?.editMode === 'assignEdit') && (
                                        <>
                                            <Divider>員工排序 (拖曳編輯)</Divider><List dataSource={(company?.employee || []) as unknown as Array<UserType & { id: string; }>}>
                                            <div className="list-group w-100">
                                                <List.RenderRow>
                                                    {({ childrenDIVProps, item, index }: listType.itemChildProps & { item: UserType; }) => (
                                                        <SelectUI.Item
                                                            itemKey={item?.id}
                                                            multipleChecked={false}
                                                            defaultValue={0}
                                                        >
                                                            {(v) => (
                                                                <a
                                                                    onClick={v?.onClick}
                                                                    {...childrenDIVProps}
                                                                    className={`
                                                                        list-group-item
                                                                        list-group-item-action
                                                                        d-flex
                                                                        mb-2
                                                                        ${v?.isChecked(item?.id) ? 'bg-danger' : ''}
                                                                    `}
                                                                >
                                                                    <Avatar style={{ fontSize: '0.5rem' }} icon={v?.isChecked(item?.id) ? '禁止' : index + 1} />
                                                                    <span className="mx-3">
                                                                        姓名 : {item?.UserName || ''}<br />
                                                                        順位 : {index + 1}
                                                                    </span>
                                                                </a>
                                                            )}
                                                        </SelectUI.Item>
                                                    )}
                                                </List.RenderRow>
                                            </div>
                                        </List>
                                    </>
                                )
                            }
                        </Form.Status>
                    </SelectUI.Group>
                </SelectUI>
            </Form>
            <Divider />
            <div className="w-100 d-flex justify-content-end">
                <button className="btn btn-outline-primary">取消</button>
                <button className="btn btn-primary mx-2">儲存</button>
            </div>
        </>
    )
}
export default Index;