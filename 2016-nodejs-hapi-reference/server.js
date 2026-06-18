'use strict';
/* eslint global-require:0 */
process.chdir(__dirname); // consistent directory

const Hapi = require('hapi');
const chalk = require('chalk');

console.log(chalk.yellow('Starting hapijs server...')); // eslint-disable-line

const settings = {
  app: {
    baseDir: __dirname
  }
};
const server = new Hapi.Server(settings);

server.connection({
  host: 'localhost',
  port: 9006
});

require('./plugins/load.js')(server);
require('./lib/logging')(server);

// Start the server
server.start(err => {
  if (err) {
    throw err;
  }
  console.log(chalk.green('Server running at:'), chalk.cyan(server.info.uri)); // eslint-disable-line
});
