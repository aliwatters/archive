var sf = require('../lib/simple-flow.js');
var fs = require('fs');

sf.series({
	'data-1': function(next) { fs.readFile('data-1.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-3': function(next) { fs.readFile('data-3.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-2': function(next) { fs.readFile('data-2.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-x': function(next) { fs.readFile('data-x.txt','utf8', function(err, txt) { next(err, txt);}) }
}, function(errors, results) {
	console.log('Series errors:',errors);
	console.log('Series results:',results);
});

sf.parallel({
	'data-1': function(next) { fs.readFile('data-1.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-3': function(next) { fs.readFile('data-3.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-2': function(next) { fs.readFile('data-2.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-x': function(next) { fs.readFile('data-x.txt','utf8', function(err, txt) { next(err, txt);}) }
}, function(errors, results) {
	console.log('Parallel errors:',errors);
	console.log('Parallel results:',results);
});

sf.limited(2, {
	'data-1': function(next) { fs.readFile('data-1.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-3': function(next) { fs.readFile('data-3.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-2': function(next) { fs.readFile('data-2.txt','utf8', function(err, txt) { next(err, txt);}) },
	'data-x': function(next) { fs.readFile('data-x.txt','utf8', function(err, txt) { next(err, txt);}) }
}, function(errors, results) {
	console.log('Limited errors:',errors);
	console.log('Limited results:',results);
});
