import React from "react";

const SelfSettingPage = ():JSX.Element => {
    return(
        <>
            <div className={styles.selfSettingBlock}>
                <div className={styles.selfSettingHead}>
                    <div>name</div>
                </div>
                <div className={styles.selfSettingBody}>
                    <div>
                        <div>帳號</div>
                        <div>a00001</div>
                    </div>
                    <div>
                        <div>密碼</div>
                        <div>L12345678</div>
                    </div>
                    <div>
                        <div>組別/部門</div>
                        <div>xxx</div>
                    </div>
                    <div>
                        <div>公司編號</div>
                        <div>70864738</div>
                    </div>
                    <div>
                        <div>公司名稱</div>
                        <div>xxx股份有限公司</div>
                    </div>
                    <div>
                        登出
                    </div>
                </div>
            </div>
        </>
    )
}
export default SelfSettingPage;
