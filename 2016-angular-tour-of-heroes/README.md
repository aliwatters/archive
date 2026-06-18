# Angular2: Tutorial: Tour of Heroes

from: https://angular.io/docs/ts/latest/tutorial/

## Notes

At time of writing (9/25/2016) the tutorial is bug free, everything works. If you're following the tutorial and hit errors in the browser - try;

* Backtrack, double check every step - invariably you will have missed a `import foo to some.component.ts` step. 
* Restart the `npm start` script - I found the transpiling to be laggy, probably to the number of saves I make while coding. A backlog would build up and the state of the browser and code would get out of step.
* Harder bugs in my code included;
  * missing `implements OnInit` on the class, and then not precisely matching the interface (eg `ngOninit` vs `ngOnInit`), the typescript compiler won't pick up the typo.
 
