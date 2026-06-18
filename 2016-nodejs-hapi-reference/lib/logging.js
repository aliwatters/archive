'use strict';
/* eslint no-console:0 */
// logging setup for hapijs

const moment = require('moment');
const chalk = require('chalk');

const levels = {
  log: 'green',
  info: 'blue',
  warn: 'yellow',
  error: 'red'
};

const meta = level => {
  const out = moment.utc().format('YYYY-MM-DD HH:mm:ss') + ' - ';
  if (process.env.NODE_ENV === 'dev') { // eslint-disable-line
    return out + chalk[levels[level]](level) + ':';
  }
  return out + level + ':';
};

const loggable = request => {
  // don't log status url or OPTIONS requests
  if (request.url.path === '/server/status' || request.method === 'options') {
    return false;
  }
  return true;
};

const setup = server => {
  server.ext('onRequest', (request, reply) => {
    request.pre.perf = {
      start: Number(new Date())
    };
    reply.continue();
  });

  server.on('response', request => {
    if (!loggable(request)) return;

    let level = 'info';
    if (request.response.statusCode === 404) {
      level = 'warn';
    } else if (request.response.statusCode === 500) {
      level = 'error';
    }
    // record timings
    request.pre.perf.end = Number(new Date());
    console[level](meta(level), request.method.toUpperCase(), request.url.path,
      String(request.pre.perf.end - request.pre.perf.start) + 'ms',
      request.response.statusCode || 0);
  });
};

module.exports = setup;
