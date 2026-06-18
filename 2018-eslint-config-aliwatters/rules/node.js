module.exports = {
  'env': {
    'node': true
  },

  'rules': {
    // enforce return after a callback
    'callback-return': 2,

    // require all requires be top-level
    // http://eslint.org/docs/rules/global-require
    'global-require': 2,

    // enforces error handling in callbacks (node environment)
    'handle-callback-err': 2,

    // disallow mixing regular variable and require declarations
    "no-mixed-requires": [
      2,
      {
        "grouping": true,
        "allowCall": true
      }
    ],

    // disallow use of new operator with the require function
    'no-new-require': 2,

    // disallow string concatenation with __dirname and __filename
    'no-path-concat': 2,

    // disallow use of process.env
    'no-process-env': 2,

    // disallow process.exit()
    'no-process-exit': 2,

    // disallow use of synchronous methods (off by default)
    'no-sync': 2
  }
};
