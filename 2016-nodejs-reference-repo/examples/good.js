'use strict';
// passes rules both jshint and jscs

function good() {
  // correctly indented - 2 spaces
  var isGood = true; // hoisted
  if (isGood === true) { // triple eq
    console.log('everything is awesome!'); // single quoted
  }
}

good();
