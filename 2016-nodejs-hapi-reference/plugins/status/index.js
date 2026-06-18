'use strict';

// const path = require('path');
// const appDir = path.dirname(require.main.filename);

exports.register = (server, options, next) => {
  // Add the route
  server.route({
    method: 'GET',
    path: '/server/status',
    handler: (request, reply) => reply('OK')
  });

  return next();
};

exports.register.attributes = {
  pkg: require('./package.json') // eslint-disable-line
};
