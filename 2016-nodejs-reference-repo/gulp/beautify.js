'use strict';

const prettify = require('gulp-jsbeautifier');
const diff = require('gulp-diff').diff;
const diffReporter = require('gulp-diff').reporter;

module.exports = (gulp, conf) => {
  gulp.task('beautify', () => {
    const task = gulp.src(conf.build.get('/lint'))
      .pipe(prettify({
        config: '.jsbeautifyrc',
        mode: 'VERIFY_AND_WRITE'
      }))
      .pipe(diff())
      .pipe(diffReporter());
    if (conf.args.write) {
      // if task is run with `--write` then overwrite source files
      task.pipe(gulp.dest('.'));
    } else {
      task.on('data', data => {
        if (data.diff && Object.keys(data.diff).length) {
          // record that there have been errors
          gulp.fail = true;
        }
      });
    }
    return task;
  });
};
