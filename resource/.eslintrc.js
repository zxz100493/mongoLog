module.exports = {
  root: true,

  env: {
    node: true
  },

  extends: [
    'plugin:vue/vue3-essential',
    '@vue/standard',
    '@vue/typescript/recommended'
  ],

  parserOptions: {
    ecmaVersion: 2020,
    parser: '@typescript-eslint/parser'
  },

  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    "no-unused-vars": "off"
  },

  extends: [
    'plugin:vue/vue3-essential',
    '@vue/standard',
    '@vue/typescript/recommended',
    '@vue/typescript'
  ]
}
