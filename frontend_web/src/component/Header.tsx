import * as React from 'react'
import { useLocation, useNavigate } from 'react-router-dom'

const Header = (): JSX.Element => {
    const navigate = useNavigate()
    const { pathname, key } = useLocation()
    console.log(pathname, key)
    return (
        <>
            <div className={styles.headerBlock}>
            </div>
            <div className={styles.navBarBlock}>
                <div
                    className={
                        pathname === '/edit'
                            ? styles.navFocus
                            : styles.navBlur
                    }
                    onClick={() => navigate('edit')}>
                    排班
                </div>
                <div
                    className={
                        pathname === '/listSetting'
                            ? styles.navFocus
                            : styles.navBlur
                    }
                    onClick={() => navigate('listSetting')}>
                    班表設定
                </div>
                <div
                    className={
                        pathname === '/listSearch'
                            ? styles.navFocus
                            : styles.navBlur
                    }
                    onClick={() => navigate('listSearch')}>
                    查看班表
                </div>
                <div
                    className={
                        pathname === '/selfSetting'
                            ? styles.navFocus
                            : styles.navBlur
                    }
                    onClick={() => navigate('selfSetting')}>
                    個人
                </div>
            </div>
        </>
    )
}
export default Header
