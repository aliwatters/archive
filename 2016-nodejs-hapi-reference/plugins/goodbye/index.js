'use strict';

// const path = require('path');
// const appDir = path.dirname(require.main.filename);
// Here load shared libraries - require(appDir + 'lib/blah')

exports.register = (server, options, next) => {
  // Add the route
  server.route({
    method: 'GET',
    path: '/goodbye',
    handler: (request, reply) => reply('goodbye world')
  });

  next();
};

exports.register.attributes = {
  pkg: require('./package.json') // eslint-disable-line
};
