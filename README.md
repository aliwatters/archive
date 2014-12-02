```
$ npm i  // installs private repo
$ npm shrinkwrap // creates npm-shrinkwrap.json
$ rm -r node_modules
$ npm i // blows up
```

Basically - npm is checking versions on the public npm module named "confit" - finds that the version 0.0.16 is not found and errors.

Required behavior - if the source is non-npm registry - don't download info from the registry about the package.

Work-around - after ```npm shrinkwrap``` modify the version of the private dependency to "*" and npm i
