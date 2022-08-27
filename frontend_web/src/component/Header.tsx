import * as React from 'react';
import styles from '../index.module.scss';
import { useNavigate } from 'react-router-dom';
const Header = () => {
    const navigate = useNavigate();
    return (
        <>
            <div className={styles.headerBlock}>
            </div>
            <div className={styles.navBarBlock}>
                <div onClick={() => navigate('edit')}>排班</div>
                <div>組別設定</div>
                <div>navOption 3</div>
                <div>navOption 4</div>
                <div>navOption 5</div>
            </div>
        </>
    )
}
export default Header;
