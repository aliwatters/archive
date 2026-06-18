function example() { // on newline - we don't want this
  var bad = false; // not hoisted
  // doesn't conform to coding standards or style!
  if (bad == false) { // only double eq
    // bad indentation
    console.log('everything is awful!'); // double quoted - we want single
  }

}

example(); // and no trailing line on EOF
