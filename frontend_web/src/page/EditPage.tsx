import React from "react";
const data = [
    {
        groupName: '保育組',
        groupMember: [
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            }
        ]
    },
    {
        groupName: '公關組',
        groupMember: [
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            }
        ]
    },
    {
        groupName: '行政組',
        groupMember: [
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            },
            {
                name: 'json',
                id: 9230994802,
            }
        ]
    }
]
const EditPage = ():JSX.Element => {
    return (
        <>  
            {data.map((dataItem) => (
                
                <div className={styles.editBlock}>
                    <div className={styles.editHeader}>{dataItem.groupName}</div>

                    <div className={styles.editBody}>
                        {
                            dataItem.groupMember.map((pItem, pIndex) => (
                                <div className={styles.editPerson}>
                                    {
                                        new Array(32).fill('').map((item, dIndex) => (
                                            <div className={styles.editDay}>
                                                {
                                                dIndex === 0
                                                    ? pItem.name
                                                    :pIndex + '-' + dIndex
                                                }
                                            </div>
                                        ))
                                    }
                                </div>
                            ))
                        }
                    </div>
                </div>
            ))
            }
        </>
    )
}
export default EditPage;