var assert = require("assert"), SF = require("../lib/simple-flow.js"), sf = new SF();


describe('Simple Flow', function() {
	
	describe('sf.series([callbacks],final-callback)', function() {


		it('should have [1,2,3] in hits when run in series', function(done) {
				var hits = [];
				sf.series([
							function(cb) { hits.push(1); cb() },
							function(cb) { hits.push(2); cb() },
							function(cb) { hits.push(3); cb() }
						],
					function(errors,results) { 
						for(err in errors) if (typeof err == 'Error') throw err; 
						assert(hits, [1,2,3]);
						done();
					}
				)
		});

		it('should have [1,2,3] for results when 3 callbacks return 1,2,3 with no errors', function(done) {
				sf.series([
							function(cb) { cb(null,1); },
							function(cb) { cb(null,2); },
							function(cb) { cb(null,3); }
						],
					function(errors,results) { 
						for(err in errors) if (typeof err == 'Error') throw err; 
						assert(results, [1,2,3]);
						done();
					}
				)
		});

		it('should have [1,2,3] for results when reading 3 text files (1,2,3) with no errors', function(done) {
				var fs = require('fs');
				sf.series([
							function(cb) { fs.readFile('data-1.txt', function(err,txt) { cb(err,txt); }) },
							function(cb) { fs.readFile('data-2.txt', function(err,txt) { cb(err,txt); }) },
							function(cb) { fs.readFile('data-3.txt', function(err,txt) { cb(err,txt); }) },
						],
					function(errors,results) { 
						for(err in errors) if (typeof err == 'Error') throw err; 
						assert(results, ['1','2','3']);
						done();
					}
				)
		});

		it('should have [1,undefined,undefined] for results when reading 3 text files (1,2,3) with error on 2', function(done) {
				var fs = require('fs');
				sf.series([
							function(cb) { fs.readFile('data-1.txt', function(err,txt) { cb(err,txt); }) },
							function(cb) { fs.readFile('data-x.txt', function(err,txt) { cb(null,txt); }) },
							function(cb) { fs.readFile('data-3.txt', function(err,txt) { cb(err,txt); }) },
						],
					function(errors,results) { 
						for(err in errors) if (typeof err == 'Error') throw err; 
						assert(results, ['1', undefined, undefined]);
						done();
					}
				)
		});

		it('should have [1,undefined,3] for results when reading 3 text files (1,2,3) with error on 2 in parallel', function(done) {
				var fs = require('fs');
				sf.parallel([
							function(cb) { fs.readFile('data-1.txt', function(err,txt) { cb(err,txt); }) },
							function(cb) { fs.readFile('data-x.txt', function(err,txt) { cb(null,txt); }) },
							function(cb) { fs.readFile('data-3.txt', function(err,txt) { cb(err,txt); }) },
						],
					function(errors,results) { 
						for(err in errors) if (typeof err == 'Error') throw err; 
						assert(results, ['1', undefined, '3']);
						done();
					}
				)
		});

		it('should have [1,undefined,3] for results when reading 3 text files (1,2,3) with error on 2 in limited', function(done) {
				var fs = require('fs');
				sf.limited(2, [
							function(cb) { fs.readFile('data-1.txt', function(err,txt) { cb(err,txt); }) },
							function(cb) { fs.readFile('data-x.txt', function(err,txt) { cb(null,txt); }) },
							function(cb) { fs.readFile('data-3.txt', function(err,txt) { cb(err,txt); }) },
						],
					function(errors,results) { 
						for(err in errors) if (typeof err == 'Error') throw err; 
						assert(results, ['1', undefined, '3']);
						done();
					}
				)
		});

	}) 
});
