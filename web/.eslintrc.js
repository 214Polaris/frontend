module.exports = {
  "env": {
    "browser": true,
    "es2021": true,
    "node": true,
  },
  "parser": "vue-eslint-parser",
  "extends": [
    "eslint:recommended",
    "plugin:vue/vue3-recommended",
    "plugin:@typescript-eslint/recommended"
  ],
  "parserOptions": {
    "ecmaVersion": 12,
    "parser": "@typescript-eslint/parser",
    "sourceType": "module",
    "requireConfigFile": false
  },
  "plugins": [
    "vue3",
    "@typescript-eslint"
  ],
  "rules": {
  }
};
