/* Module simple-flow.js -- simple model
 * Based on examples in Mixu's node book http://book.mixu.net/ 
 * added basic error handling, named callbacks, results.
 * Limitation -- doesn't handle throw in callbacks, does anything?
 */

module.exports.series = function(callbacks, last) {
	var context = {results: [], errors : [], keys: [], 'callbacks' : callbacks}, index = 0;
	function next(context) {
		var callback = context.callbacks.shift();
		if (callback) {
			callback(function(err,res) {
				if (context.keys.length > 0) {
					context.errors[context.keys[index]] = err;
					context.results[context.keys[index]] = res;
				} else {
					context.errors[index] = err;
					context.results[index] = res;
				}
				index++;
				// as series we care about errors -- end on error.
				// if you don't care about errors -- use parallel or limited -- or pass null
				if (err) last(context.errors, context.results)
				else next(context);
			});
		} else {
			last(context.errors, context.results);
		}
	}
	
	prep_callbacks(context);
	next(context);
}


module.exports.parallel = function(callbacks, last) {
	var context = {results:[],errors:[],keys:[],'callbacks':callbacks}, result_count = 0;

	prep_callbacks(context);

	context.callbacks.forEach( function(callback, index) {
		callback( function(err, res) {
			if (context.keys.length > 0) {
				context.errors[context.keys[index]] = err;
				context.results[context.keys[index]] = res;
			} else {
				context.errors[index] = err;
				context.results[index] = res;
			}
			result_count++;
			if (result_count == context.callbacks.length) {
				last(context.errors, context.results);
			}
		});
	});
}



module.exports.limited = function(limit, callbacks, last) {
	var context = {results:[],errors:[],keys:[],'callbacks':callbacks},
		running = 1, task = 0;
	
	function runLimited(context,index) {
		context.callbacks[index](function(err, res) {
			if (context.keys.length > 0) {
				context.errors[context.keys[index]] = err;
				context.results[context.keys[index]] = res;
			} else {
				context.errors[index] = err;
				context.results[index] = res;
			}
			next(context);
		});
	}
	
	function next(context) {
		running--;
		if (task === context.callbacks.length && running === 0) {
			last(context.errors,context.results);
		}
		while (running < limit && context.callbacks[task]) {
			runLimited(context,task++);
			running++;
		}
	}
	
	prep_callbacks(context);
	next(context);
}



function prep_callbacks(context) { // should work on calling functions scope.
	if (Array.isArray(context.callbacks)) {
		// default behavior	based on Array of functions
	} else if ( typeof context.callbacks == 'object') {
		context.keys = Object.keys(context.callbacks);
		var values = [];
		for (var k in context.callbacks) values.push(context.callbacks[k]);
		context.callbacks = values;
		// is values GC'd?
	} else {
		// TBA -- throwing an error -- should this be done?
		throw new Error('usage: sf.parallel([callbacks],final-callback) or sf.parallel({ name1 : function , name2 : function }, final-callback)');
	}
}
