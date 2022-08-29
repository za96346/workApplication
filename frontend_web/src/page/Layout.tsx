import React from "react";
import { Outlet } from "react-router-dom";
import Footer from "../component/Footer";
import Header from "../component/Header";

const Layout = ():JSX.Element => {
    return (
        <>
        <Header />
            <div className={styles.article}>
                <Outlet />
            </div>
        <Footer />
        </>
    )
}
export default Layout;