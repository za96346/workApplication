import React from "react";
import styles from '../index.module.scss';
const EditPage = () => {
    return (
        <div className={styles.editBlock}>
            <div className={styles.editGroup}>
                <div className={styles.editHeader}>xx 組</div>

                <div className={styles.editBody}>
                    {
                        new Array(5).fill('').map(() => (
                            <div className={styles.editPerson}>
                                <div className={styles.editDay}>
                                    
                                </div>
                            </div>
                        ))
                    }
                </div>
            </div>
        </div>
    )
}
export default EditPage;