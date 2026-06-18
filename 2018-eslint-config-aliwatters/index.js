module.exports = {
  extends: [
    './rules/best-practices',
    './rules/errors',
    './rules/es6',
    './rules/imports',
    './rules/node',
    './rules/style',
    './rules/strict',
    './rules/variables'
  ].map(require.resolve),
  env: {
    amd: false,
    browser: false,
    node: true,
    jasmine: false,
    mocha: true
  },
  parserOptions: {
    ecmaVersion: 6,
    sourceType: 'script',
  }
};
