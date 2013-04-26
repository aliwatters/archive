var SF = require('./lib/simple-flow.js');
var sf = new SF;
var fs = require('fs');


sf.parallel([
	function(next) { fs.readFile('data-1.txt','utf8', function(err, txt) { next(err, txt);}) },
/*	function(next) { fs.readFile('data-x.txt','utf8', 
							function(err, txt) {
								 if (err) throw err; // simulating a library throwing an error
								 next(err, txt);
							}) 
	}, // Note -- unforunately -- a thrown error cannot be caught -- tried with try catch and with domains
*/
	function(next) { fs.readFile('data-x.txt','utf8', function(err, txt) { next(err, txt);}) },
	function(next) { fs.readFile('data-3.txt','utf8', function(err, txt) { next(err, txt);}) },
	function(next) { fs.readFile('data-2.txt','utf8', function(err, txt) { next(err, txt);}) }
], function(errors, results) {
	console.log('Parallel errors:',errors);
	console.log('Parallel errors:',results);
});


sf.series([
	function(cb) { fs.readFile('data-1.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-2.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-3.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-x.txt','utf8', function(err, txt) { cb(err, txt);}) }
], function(errors, results) {
	console.log('Series errors',errors);
	console.log('Series results',results);
});

// Ignore error series example

sf.series([
	function(cb) { fs.readFile('data-1.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-x.txt','utf8', function(err, txt) { cb(null, txt);}) }, // I don't care about err here
	function(cb) { fs.readFile('data-2.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-3.txt','utf8', function(err, txt) { cb(err, txt);}) }
], function(errors, results) {
	console.log('Series 2 errors',errors);
	console.log('Series 2 results',results);
});

// Limited

sf.limited(2, [
	function(cb) { fs.readFile('data-1.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-x.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-2.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-3.txt','utf8', function(err, txt) { cb(err, txt);}) }
], function(errors, results) {
	console.log('Limited errors',errors);
	console.log('Limited results',results);
});



// Mixed example...
// working! -- name callbacks different in the sub -- stuff, and then in the result callback to the original context

sf.series([
	function(cb) { 
		sf.parallel([
			function (cb1) { fs.readFile('data-1.txt','utf8', function(err, txt) { cb1(err, txt);}) },
			function (cb1) { fs.readFile('data-2.txt','utf8', function(err, txt) { cb1(err, txt);}) },
			function (cb1) { fs.readFile('data-x.txt','utf8', function(err, txt) { cb1(err, txt);}) }
		] ,function (errors, results) {
				console.log('Sub Errors:',errors);
				console.log('Sub Results:',results);
				cb(null,results); // error handling here? -- if we add a err instead of null -- series dies.
			}
		)
	},
	function(cb) { fs.readFile('data-3.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-x.txt','utf8', function(err, txt) { cb(err, txt);}) }
], function(errors, results) {
	console.log('Mixed errors',errors);
	console.log('Mixed results',results);
});



