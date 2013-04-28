simple-flow
===========

Node.js simple control flow library, implements series, parallel and limited models. Returns an array of errors and an array of results.

Initially based on examples presented in http://book.mixu.net/ -- read this, great way to get going in node.


## Usage: See the examples directory for more.

```javascript
var sf = require('./lib/simple-flow.js');

// General use
sf.series([callbacks], final);
sf.parallel([callbacks], final);
sf.limited(limit, [callbacks], final);
```

## Series method

```javascript
// For simple-flow examples we use fs: note -- data-x.txt is deliberatly non-existant
var fs = require('fs');

// sf.series([callbacks], final-callback)

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
```

## Parallel operations

```javascript
// sf.parallel([callbacks], final-callback)

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
```


# Limited operation -- eg a pool of database clients that you don't want to flood

```javascript
// sf.limited(int, [callbacks], final-callback)

sf.limited(2, [
	function(cb) { fs.readFile('data-1.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-x.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-2.txt','utf8', function(err, txt) { cb(err, txt);}) },
	function(cb) { fs.readFile('data-3.txt','utf8', function(err, txt) { cb(err, txt);}) }
], function(errors, results) {
	console.log('Limited errors',errors);
	console.log('Limited results',results);
});
```

# Mixed operation -- starts to get complex!

```javascript
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
```



## Limitations

Can't catch errors in callbacks that are thrown.


## Future expansion

Do named errors and results as done in async
