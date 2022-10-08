/* eslint-disable no-restricted-globals */
/* eslint-disable class-methods-use-this */

export const placeholder = {
  inputPhone: '請輸入電話號碼',
  inputPhoneCaptcha: '請輸入手機簡訊驗證碼',
  inputPassword: '請輸入密碼',
  confirmPassword: '確認密碼',
  captchaUpperOrLowerCase: '驗證碼',
  promoteCode: '邀請碼(非必填)',
  checked: '尚未同意',
};
class Rule {
  private rules = {
    onlyNum: { required: true, message: '只允許數字', pattern: /^[0-9\s]*$/ },
  };

  isPhone(value: string | number): boolean {
    if (!/^[0-9\s]*$/.test(`${value}`)) {
      return true;
    }
    return false;
  }

  public email() {
    return [
      { required: true, message: '不允許空白!', whitespace: true },
    ];
  }

  public password() {
    return [
      { required: true, message: '非法字元!', pattern: /^[a-zA-Z0-9 ]+$/ },
      { required: true, message: '不允許空白!', whitespace: true },
      { required: true, message: '請輸入至少八碼!', min: 8 },
    ];
  }

  public passwordConfirm() {
    return [{ required: true, message: placeholder.confirmPassword }, ...this.password()];
  }

  public captcha() {
    return [{ required: true, message: '請輸入六碼', len: 6 }, this.rules.onlyNum];
  }

}
export default new Rule();
