/*
JS

- Loosely-typed language
- Functional
- Interpreted at runtime (V8 / node.js)
- Object
- Prototype
- Class (ES6)
*/

var i = 0;
let j = 0;

j = "foo"; // string

// Javascript with parentheses and semicolons
for (let i = 0; i < 10; i++) {
    console.log(i);
}

// Javascript assignment
var foo = "bar";


let foo = "bar";


bar = "blah";



// Javascript export
const Bar = () => {};

module.exports = {
  Bar
}


// Javascript import
var foo = require('foo');
foo.bar();



// Javascript - return multiple values
function foo() {
    return {a: 1, b: 2};
}
const { a, b } = foo();


// Node error handling
foo('bar', function(err, data) {
    //handle error

    if (err) console.log(err);
});


// Variadic function
function foo (...args) {
    console.log(args.length);
}

foo(); // 0
foo(1, 2, 3); // 3