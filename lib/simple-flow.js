/* Module simple-flow.js -- simple model
 * Based on examples in Mixu's node book http://book.mixu.net/ 
 * added basic error handling.
 * Limitation -- doesn't handle throw in callbacks, does anything?
 * Near future plans -- extend to do named callbacks eg. { 'name': function, 'name2': function }
 */

module.exports.series = function(callbacks, last) {
	var results = [], errors = [];
	function next() {
		var callback = callbacks.shift();
		if (callback) {
			callback(function(err,res) {
				errors.push(err);
				results.push(res);
				// as series we care about errors -- end on error.
				// if you don't care about errors -- use parallel or limited -- or pass null
				if (err) last(errors, results)
				else next();
			});
		} else {
			last(errors, results);
		}
	}
	
	next();
}


module.exports.parallel = function(callbacks, last) {
	var results = [], errors = [], result_count = 0;

	callbacks.forEach( function(callback, index) {
		callback( function(err, res) {
			errors[index] = err;
			results[index] = res;
			result_count++;
			if (result_count == callbacks.length) {
				last(errors, results);
			}
		});
	});
}


module.exports.limited = function(limit, callbacks, last) {
	var results = [], errors = [],
		running = 1, task = 0;
	
	function runLimited(index) {
		callbacks[index](function(err, res) {
			errors[index] = err;
			results[index] = res;
			next();
		});
	}
	
	function next() {
		running--;
		if (task === callbacks.length && running === 0) {
			last(errors,results);
		}
		while (running < limit && callbacks[task]) {
			runLimited(task++);
			running++;
		}
	}
	
	next();
}

