/* eslint no-console:0, no-process-exit:0, curly:["error", "multi"], global-require:0, no-process-env:0, no-sync:0 */
'use strict';
const eslint = require('gulp-eslint');

module.exports = (gulp, conf) => {
  gulp.task('lint', () => {
    gulp.src(conf.build.get('/lint'))
      .pipe(eslint())
      .pipe(eslint.format())
      .pipe(eslint.failOnError())
      .on('error', () => {
        // record that there have been errors
        gulp.fail = true;
      });
  });
};
