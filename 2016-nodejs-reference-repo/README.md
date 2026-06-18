# Nodejs Reference - a repo containing config for nodejs projects

Aim here is to automate as many of the linting and styling steps, and provide a good starting point for a nodejs based project.

NOTE: requires git > 1.8

Currently includes:

* gitignore - don't check in logs, node_modules, tempfiles and the like.
* editorconfig - www.editorconfig.org - use the editor of your choice and remain consistent.
* eslint - lint javascript files - don't put style rules here
* jsbeautify - fix formatting automatically

```
npm run lint
npm run beautify
```
