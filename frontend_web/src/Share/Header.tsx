import { MenuOutlined } from '@ant-design/icons'
import * as React from 'react'
import { useLocation } from 'react-router-dom'

const Header = ({ setShow }: { setShow: (arg0: boolean) => void }): JSX.Element => {
    // const navigate = useNavigate()
    const { pathname, key } = useLocation()
    console.log(pathname, key)
    return (
        <>
            <div onClick={() => setShow(true)} className={window.styles.headerBlock}>
                <MenuOutlined />
            </div>
        </>
    )
}
export default Header
