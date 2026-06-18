/* eslint no-process-exit:0, no-sync:0, no-console:0, global-require:0, prefer-template:0 */
'use strict';

const gulp = require('gulp');
const path = require('path');
const parseArgs = require('minimist');
const Confidence = require('confidence');
const conf = {
  build: new Confidence.Store(require('./conf/build')),
  args: parseArgs(process.argv.slice(2)),
  gwd: __dirname
};

const avail = [];
// combine all the gulp tasks
require('fs').readdirSync('./gulp').forEach(file => {
  if (path.extname(file) === '.js') {
    avail.push(file.substr(0, file.indexOf('.js')));
    try {
      require('./gulp/' + file)(gulp, conf);
    } catch (err) {
      console.error('ERROR LOADING:', file, err);
    }
  }
});

gulp.task('default', () => {
  console.log('Available Gulp Tasks:', avail);
});

// Aliases - short ones for rollup tasks
// gulp.task('lint', ['eslint', 'js-beautify'], function() {});

process.on('exit', () => {
  if (gulp.fail) {
    // return non-zero exit code
    process.exit(1);
  }
});
