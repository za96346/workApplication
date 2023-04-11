module.exports = {
  "root": true,
  "env": {
      "browser": true,
      "es2021": true
  },
  "extends": [
      "@react-native-community",
      "standard-with-typescript"
      // "plugin:@typescript-eslint/eslint-plugin"
  ],
  "parser": "@typescript-eslint/parser",
  "overrides": [
  ],
  // 解决tsconfig下的path别名导致eslint插件无法解决的bug
  "parserOptions": {
      "ecmaVersion": "latest",
      "sourceType": "module",
      "project": "./tsconfig.json",
      "tsconfigRootDir": ""
  },
  "plugins": [
      "react",
      "import"
  ],
  "rules": {
      "indent": "off",
      "quotes": "off",
      "new-cap": "off",
      "@typescript-eslint/indent": ["error", 4],
      "@typescript-eslint/strict-boolean-expressions": "off",
      "@typescript-eslint/quotes": "warn",
      "@typescript-eslint/no-misused-promises": "warn",
      "@typescript-eslint/restrict-template-expressions": "off",
      "@typescript-eslint/consistent-type-assertions": "off",
      "@typescript-eslint/no-floating-promises": "warn",
      "import/no-absolute-path": "off"
  },
  "settings": {
      "typescript": {
          "alwaysTryTypes": true
      }
  }
}
