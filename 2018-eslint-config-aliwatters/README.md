JavaScript ESLint config based on https://github.com/airbnb/javascript represents the personal preferences of the author - Ali Watters

Add package and eslint:

`npm install --save-dev git+ssh://git@github.com/aliwatters/eslint-config-aliwatters.git eslint eslint-plugin-import`

Add a .eslintrc file to the root of your project:

```
{
  "extends": "aliwatters"
}
```
