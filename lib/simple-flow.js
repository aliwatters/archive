// Module control-flow.js -- simple model
// Based on examples in Mixu's node book http://book.mixu.net/ 
// added basic error handling.
// Limitation -- doesn't handle throw in callbacks, does anything?
// Near future plans -- extend to do named callbacks eg. { 'name': function, 'name2': function }

function SimpleFlow() {
}

SimpleFlow.prototype.series = function(callbacks, last) {
	var results = [];
	var errors = [];
	function next() {
		var callback = callbacks.shift();
		if (callback) {
			callback(function() {
				var args = Array.prototype.slice.call(arguments);
				errors.push(args[0]);
				results.push(args[1]);
				// as series we care about errors -- end on error.
				// if you don't care about errors -- use parallel or limited -- or pass null
				if (args[0]) last(errors, results)
				else next();
			});
		} else {
			last(errors, results);
		}
	}
	
	next();
}


SimpleFlow.prototype.parallel = function(callbacks, last) {
	var results = [];
	var errors = [];
	var result_count = 0;

	callbacks.forEach( function(callback, index) {
		callback( function() {
			var args = Array.prototype.slice.call(arguments);
			errors[index] = args[0];
			results[index] = args[1];
			result_count++;
			if (result_count == callbacks.length) {
				last(errors, results);
			}
		});
	});
}


SimpleFlow.prototype.limited = function(limit, callbacks, last) {
	var results = [];
	var errors = [];
	var running = 1;
	var task = 0;
	
	function next() {
		running--;
		if (task == callbacks.length && running == 0) {
			last(errors,results);
		}
		while (running < limit && callbacks[task]) {
			var callback = callbacks[task];
			(function(index) { // anonymous function maintains orig scope
				callback(function() {
					var args = Array.prototype.slice.call(arguments);
					errors[index] = args[0];
					results[index] = args[1];
					next();
				});
			})(task);
			task++;
			running++;
		}
	}
	
	next();
}

module.exports = SimpleFlow;
