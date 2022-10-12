class Language {
    msg: any
    loginSuccess = '登入成功'
    formatError = '資料格式錯誤'
    noUser = '沒有此使用者'
    thisAccountIsNotRegisted = '此帳號還沒有被註冊 請先註冊帳號'
    accountOrPasswordError = '帳號或密碼錯誤'
    emailCaptchaSendedSuccess = '驗證碼發送成功'
    emailCaptchaSendedFail = '驗證碼發送失敗'
    registerSuccess = '註冊成功'
    accountHasBeenRegistered = '此信箱已被註冊'
    registerFail = '註冊失敗 ， 請稍候再試'
    captchaIsNotCorrect = '驗證碼不正確'
    passwordIsNotSame = '密碼不相等'
    emailIsNotCorrect = '電子信箱格式錯誤'
    tokenExpire = 'token過期'
    tokenNotAcceptable = '無效token'
    notAcceptableId = '無效的id, 請重新查詢'
}
export default new Language()
