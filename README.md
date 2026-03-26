# FLanguage

Welcome to FLanguage! This project was created as an experiment to understand how interpreters work and to learn the Go programming language. You can find an example of FLanguage code in `helloWorld.txt`.

- **RunFile**: To run FLanguage code, navigate to `./bin/` and execute `.\FLanguage.exe "<file.txt>"`
- **RunRepl**: To run the FLanguage REPL, navigate to `./bin/` and execute `.\FLanguage.exe "r"`. Use `{{` and `}}` for multi-line code.

## Key Features

- **Functions**: Define and use functions to organize your code into reusable modules. Functions cannot modify the external environment in which they are defined.
- **Arrays**: Easily manipulate data collections using arrays, allowing you to store and access elements efficiently.
- **Hash Tables**: Use hash tables to implement key-value data structures, ideal for managing data associations and inline functions.
- **Conditional Statements**: Control program flow using if-else statements, allowing different operations to be executed based on specific conditions.
- **While Loops**: Iterate through data or repeat operations as long as a specific condition is true.
- **Inline Functions**: Define functions directly within the main code context to improve readability and modularity.
- **External Code Import**: Easily import code from other files to organize your project into separate modules and reuse existing code. Modules must consist exclusively of functions. To call a function from a module, concatenate the module name (without extension) with the function name, separated by an underscore (`_`).
- **Objects**: Hash tables can contain inner functions that interact with the hash table itself through the `this` keyword.
- **End of File**: The end of every file must be indicated with the `END` keyword.

## Syntax

- **let** — Declare a variable. Once a type is assigned, it cannot be changed.
  ```
  let a = 2;
  ```

- **import** — Import a module and return it as an object.
  ```
  let tree = import("BinarySearch.txt");
  tree{"Run"}([1,2,3,4,7], 4);
  ```

- **Ff** — Declare a named function.
  ```
  Ff getMatrix() {
      ret [[2,4],[2,3,4]];
  }
  ```

- **@** — Define an anonymous inline function, assignable to variables, hash tables, or arrays.
  ```
  let a = @(a, b) {
      ret a + b;
  };
  let b = a(2, 1);
  ```

- **ret** — Return a value from a function.
  ```
  Ff greet() {
      ret "hello";
  }
  ```

- **if / else** — Conditional branching.
  ```
  if (4 < 2) {
      a = a + 2;
  } else {
      a = a * 4;
  }
  ```

- **while** — Loop while a condition is true.
  ```
  while (i < 5) {
      i = i + 1;
  }
  ```

- **this{}** — Reference the current hash table from within an inner function.
  ```
  let object = {
      "name": "luca",
      "age": 22,
      "birthday": @() {
          this{"age"} = this{"age"} + 1;
          ret this{"age"};
      }
  };
  object{"birthday"}();
  ```

## Built-in Functions

| Function | Parameters | Description |
|---|---|---|
| `len` | `a` | Returns the number of elements in an array, or the number of characters in a string |
| `newArray` | `n`, `type` | Creates an array of size `n` with every element initialized to `type` |
| `int` | `a` | Converts a string or float to an integer |
| `float` | `a` | Converts a string or integer to a float |
| `string` | `a` | Converts a value to a string |
| `print` | `a` | Prints `a` to the console |
| `println` | `a` | Prints `a` to the console with a newline |
| `read` | — | Reads a line from standard input |
| `import` | `path` | Loads a module from `path` and returns it as an object |
