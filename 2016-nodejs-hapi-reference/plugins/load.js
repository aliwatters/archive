'use strict';
/* eslint global-require:0, no-sync:0, no-console:0 */

// Loads all the plugins in sub-directories

const fs = require('fs');
const path = require('path');

const load = server => {
  const plugins = fs.readdirSync(__dirname)
    .map(a => path.join(__dirname, a))
    .filter(file => fs.statSync(file).isDirectory());

  const loaded = [];
  plugins.forEach(dir => {
    const name = path.basename(dir);
    loaded.push(name);
    const plugin = require(dir);
    server.register(plugin, err => {
      if (err) {
        console.error('Failed to load plugin:', name, err);
      }
    });
  });
  console.log('Loaded Plugins:', loaded);
};

module.exports = load;
