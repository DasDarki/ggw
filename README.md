# ggw
Just an experimental programming language for webdev.

## Specification
At heart the language is statically typed. Throughout my years of programming I have come to the conclusion that
static typing is a must for me. It prevents a lot of errors and makes the code more readable and maintainable.

### General Thoughts
So, at the point of writing this, I wanted to try to create language which just fits perfectly into the space of
web development and I wanted to try to make it as simple, powerful and versatile as possible. The language should
enable the developer to create web applications with ease and without the need to learn a lot of different languages
allowing the usage of just one language in front- and backend.

At daily basis, I use mostly TypeScript, HTML, CSS and Go. The compiler is written in Go and should compile its code
in the backend to Go code which can be compiled to a binary and run on a server. The frontend code should be compiled
to WAT code which can be run in the browser through WebAssembly.

The language itself should be expressive and will probably extremely opinionated. Lets see how this will turn out.

#### Billion Dollar Mistake
The language will not have `null` or `undefined`. I like the way Rust handles this by using `Option` and `Result`.
It forces the developer to handle the case where a value might not be present. So, language uses a similar approach.
Each value which does not have a value declared is automatically assigned its default value. This means that a
variable of type `int` is automatically assigned `0` if no value is assigned to it.

#### OOP
Its maybe an unpopular opinion but I like OOP. The language will have classes and interfaces. Classes can have
properties and methods. Interfaces can have properties and methods but no implementation. A class can implement
multiple interfaces.

There is inheritance, overloading and generics.

### Type System
The type system is inspired by TypeScript and Go. It is a structural type system which means that types are defined
by their structure and not by their name. This means that two types are equal if they have the same structure.

#### Primitive Types
- `int`: A 64-bit signed integer.
- `float`: A 64-bit floating point number.
- `bool`: A boolean value.
- `string`: A UTF-8 encoded string.

#### Generic Types
Wit `<T>` before names of types & methods, you can define generic types. Generics can be made more specific by
defining constraints. Constraints can be interfaces or classes.

#### Composite Types
- `array` -> `[]T`: An array of elements of type `T`.
- `map` -> `[T]U`: A map of keys of type `T` and values of type `U`.
- `tuple` -> `(T, U, V)`: A tuple of elements of type `T`, `U` and `V`.

#### Type Inference
The language has type inference. This means that the type of a variable can be inferred by the compiler (hopefully).

### Syntax
#### Error Handling
I must say it: Go's error handling is big pile of crap. I hate it. But the idea of having errors as values is great.
The try-catch hell is just not the way to go. So, the language will force the developer to handle errors. From the point
of writing I dont want to use panic either. The problem with it: I use it at the wrong places. An application should never panic
out of the flow. It should always be handled.

#### Comments
Comments are written with `//` and are single line comments.

```ggw
// This is a comment.
```

Comments can also be written as block comments.

```ggw
/// This is a block comment.
```

Block comments can use annotations to generate documentation.

#### Variables
Variables are declared with the `let` keyword.

```ggw
let x: int = 42;
```

Variables can be declared without a type. The type is then inferred by the compiler.

```ggw
let x = 42;
```

Variables can be declared without a value. The value is then automatically assigned its default value.

```ggw
let x: int; // x is automatically assigned 0.
```

Variables can be declared as constants with the `const` keyword. Constants cannot be reassigned.

```ggw
const x: int = 42;
```

The `let` or `const` keyword is optional for an expressive declaration. If omitted, the compiler declares it as `let` 
based on the concept "immutable by default".
```ggw
x: int = 42;
x = 43; // Error: Cannot reassign variable.
```

#### Functions
Functions are declared with the `func` keyword.

```ggw
func add(a: int, b: int): int {
    return a + b;
}
```

Functions can be declared without a return type. The return type is then inferred by the compiler.

```ggw
func add(a: int, b: int) {
    return a + b;
}
```

If a function does not return a value, the return type is `none`.

```ggw
func printHello() {
    print("Hello, World!");
}

// OR

func printHello(): none {
    print("Hello, World!");
}
```

Functions can return tuples and tuples can be spread.

```ggw
func getPoint(): (int, int) {
    return (1, 2);
}

let x, y = getPoint();
```

Functions can have default values for parameters.

```ggw
func add(a: int, b: int = 0): int {
    return a + b;
}

add(1); // 1

add(1, 2); // 3
```

When calling functions with default parameters, the parameters can be named.

```ggw
add(a = 1); // 1

add(a = 1, b = 2); // 3
```

Functions can have variadic parameters.

```ggw
func sum(args: ...int): int {
    // args is an array of integers.
}

sum(1, 2, 3, 4, 5);
```

Parameter types can be grouped. Which means that multiple parameters can have the same type. The last parameter in a group
defines the type for all parameters in the group.

```ggw
func add(a, b: int): int {
    return a + b;
}

add(1, 2);
```

Tupled return values can be named.

```ggw
func getPoint(): (x: int, y: int) {
    return 1, 2; // here the order is important
}

// OPTIONALLY USE NAMED RETURNS:
func getPoint(): (x: int, y: int) {
    return x = 1, y = 2;
}

let point = getPoint();
print(point.x); // 1
print(point.y); // 2

// OR deconstruct the tuple
let x, y = getPoint();
```

And return values can be ignored.

```ggw
func getPoint(): (int, int) {
    return 1, 2;
}

let _, y = getPoint();
```

Or using generics.

```ggw
func getPoint<T, U>(x: T, y: U): (T, U) {
    return x, y;
}

let x, y = getPoint(1, 2);
```

Another cool trick when returning multiple values is yielding them.

```ggw
func getPoint(): (int, int) {
    yield 1;
    yield 2;
}

// OR NAMED:
func getPoint(): (x: int, y: int) {
    yield x = 1;
    yield y = 2;
    
    return; // optional. If the compiler recognized all return values, it instantly returns.
}

let x, y = getPoint();
```

Yielding allows for a cleaner control flow and visually separates the return values from the rest of the function.

#### Classes
Classes are declared with the `class` keyword.

```ggw
class Point {
    x: int;
    y: int;

    func move(x: int, y: int) {
        this.x += x;
        this.y += y;
    }
}

let point = new Point();
point.x = 1;
point.y = 2;
```

Classes can have constructors. Constructors are methods directly behind the name of the class.

```ggw
class Point(x: int, y: int) {
    x: int = x; // optional, the compiler will automatically assign the value if the name is the same
    y: int = y;

    func move(x: int, y: int) {
        this.x += x;
        this.y += y;
    }
}

let point = new Point(1, 2);
```

If a classes has values in the constructor which are not assigned to properties, the compiler will automatically create
let properties with the same name.

```ggw
class Point(x: int, y: int, z: int) {
    
    func move(ox: int, oy: int) {
        this.x += ox;
        this.y += oy;
    }
}
```

Classes can have static methods.
Static functions can be called without an instance of the class.

```ggw
class Math {
    static func add(a: int, b: int): int {
        return a + b;
    }
}
```

If you dont like the constructor syntax, you can overwrite the constructor method.

```ggw
class Point {
    x: int;
    y: int;

    constructor(x: int, y: int) {
        this.x = x;
        this.y = y;
    }
}
```

#### Inheritance
Classes can inherit from other classes.

```ggw
class Shape(x: int, y: int) {}

class Rectangle(x: int, y: int, width: int, height: int)
    extends Shape(x, y) {
    
    width: int;
    height: int;
}
```

Abstract classes can be declared with the `abstract` keyword.

```ggw
abstract class Shape {
    x: int;
    y: int;

    abstract func area(): float;
}
```

Abstract methods must be implemented by the inheriting class.

```ggw
class Rectangle(x: int, y: int, width: int, height: int)
    extends Shape {
    
    func area(): float {
        return this.width * this.height;
    }
}
```

When extending a class, you can access the parent class with the `super` keyword.

```ggw
class Rectangle(x: int, y: int, width: int, height: int)
    extends Shape {
    
    func area(): float {
        return super.area(); // i know this is illegal syntax because its an abstract class but hey its just an example
    }
}
```

And classes can also use generics.

```ggw
class Pair<T, U> {
    first: T;
    second: U;
}
```

#### Overloading
Functions can be overloaded.

```ggw
func add(a: int, b: int): int {
    return a + b;
}

func add(a: float, b: float): float {
    return a + b;
}
```

#### Interfaces
Interfaces are declared with the `interface` keyword.

```ggw
interface Shape {
    x: int;
    y: int;

    area(): float;
}
```

Classes can implement the interfaces.

```ggw
class Rectangle(x: int, y: int, width: int, height: int)
    implements Shape {
    
    func area(): float {
        return this.width * this.height;
    }
}
```

Interfaces can extend other interfaces.

```ggw
interface Shape {
    x: int;
    y: int;

    area(): float;
}

interface Drawable
    extends Shape {
    
    draw(): none;
}
```

#### Enums
Enums are declared with the `enum` keyword.

```ggw
enum Color {
    Red,
    Green,
    Blue
}
```

By default enums are zero-based, incremented number values.

```ggw
enum Color {
    Red, // 0
    Green, // 1
    Blue // 2
}
```

But they can have more complex values.

```ggw
enum Color(hex: string, value: int) {
    Red("#FF0000", 1), // order is important
    Green("#00FF00", 2),
    Blue("#0000FF", 3)
}

let color = Color.Red;
print(color.hex); // #FF0000
```

#### Access Modifiers
Properties, methods and classes can have access modifiers.
- `public`: Accessible from everywhere.
- `protected`: Accessible from the class and its subclasses.
- `private`: Accessible only from the class.

```ggw
class Point {
    public x: int;
    protected y: int;
    private z: int;
}
```

#### Control Flow
The language has the usual control flow statements.

##### If
```ggw
if x == 0 {
    print("x is zero.");
} else if x == 1 {
    print("x is one.");
} else {
    print("x is neither zero nor one.");
}
```

Syntactic sugar for if statements is the `when` statement.

```ggw
let message = when x {
    0 => "x is zero.";
    1 => "x is one.";
    else => "x is neither zero nor one.";
}
```

Or for unary in short:

```ggw
let message = x == 0 ? "x is zero." : "x is not zero.";
```

##### For
```ggw
for i = 0; i < 10; i++ {
    print(i);
}
```

For slices:

```ggw
for i, value in values {
    print(i, value);
}
```

If the index is not needed:

```ggw
for value in values {
    print(value);
}
```

And for maps:

```ggw
for key, value in values {
    print(key, value);
}
```

But also here index if needed:

```ggw
for i, key, value in values {
    print(i, key, value);
}
```

And if you have a tuple array you can instantly deconstruct:

```ggw
for (key, value) in values {
    print(key, value);
}
```

##### While
```ggw
while x < 10 {
    x++;
}
```

##### Do-While
```ggw
do {
    x++;
} while x < 10;
```

#### Modules
Modules are declared with the `module` keyword.
A module is a collection of classes, interfaces, functions and variables. It is a way to organize code.

For the ones coming from Go, modules are similar to packages.

```ggw
module math; // somewhere in the code. At best at the beginning of the file.
```

If no module is declared, the code is in the anonymous module and cant be used inside other modules. Its for internal use only.

##### Directory-based Modules
Adding a .mod file to the directory will make the directory a module. The .mod file needs to contain the module name
and can optionally declare specific directives which all files in that directory will inherit.

```ggw
module math;
```

##### Importing Modules
Modules can be imported with the `import` keyword.

```ggw
import math;
```

The compiler tries its best to resolve the module. If the module is not found, the compiler will throw an error.
If there are multiple modules with the same name you can specify the path.

```ggw
import math from "github.com/user/math";
```

Renaming modules is also possible.

```ggw
import math as m from "github.com/user/math";
```

#### Annotations
Annotations are a way to add metadata to classes, interfaces, methods and properties.
Annotations are declared with the `@` symbol.

```ggw
@deprecated
class Point {
    x: int;
    y: int;
}
```

Annotations can have values.

```ggw
@deprecated("Use the Point2 class instead.")
class Point {
    x: int;
    y: int;
}
```

Annotations should be easily created and accessible.

```ggw
annotation deprecated(message: string);
```

#### Reflection
Reflection is a way to inspect and modify code at runtime.
Reflection is a powerful tool but should be used with caution.
So it will be used declaratively.

```ggw
class Point {
    x: int;
    y: int;
}

let point = new Point();
let type = point.#reflect(TYPE);

print(type.name); // Point

let properties = type.#reflect(PROPERTIES);
for property in properties {
    print(property.name);
}
```

And with annotations:

```ggw
class Point {
    @deprecated("Use the Y instead.")
    x: int;
    y: int;
}

let point = new Point();
let property = point.#reflect(PROPERTIES).find((property) -> property.name == "x");
let annotation = property.#reflect(ANNOTATIONS).find((annotation) -> annotation.name == "deprecated");
print(annotation.message); // Use the Y instead.
```

#### And now the fun part and the most important part: WebDev
The language is intended to be used for web development. This means that the language should have a way to interact with
frontend and backend without a big hustle. I said the language is optionated which means that the language will have
a built-in way to do somethings I deem important in web development.

Here is an overview:
- `html`: A way to interact with the DOM.
- `css`: A way to style the DOM.
- `communication`: A way to communicate between frontend and backend. At best with multiple adapters (WebSockets, HTTP, etc.).
- `database`: A way to interact with databases.
- `database/ORM`: A way to interact with databases in an object-oriented way.
- `DI`: A way to handle dependency injection.
- `routing`: A way to handle routing in the frontend and backend.
- `storage`: A way to store data in the frontend and backend.
- `validation`: A way to validate data in the frontend and backend.
- `testing`: A way to test the frontend and backend.
- `logging`: A way to log in the frontend and backend.
- `security`: A way to secure the frontend and backend.

And this is just a small overview. The language should be as versatile as possible. But I should not overshoot the goal.

##### Communication
So the idea is so called wiring. You write the whole full-stack application into one project and the compiler
will wire the frontend and backend together. The frontend will be compiled to WAT code and the backend to Go code.

First we need to tell the compiler which file is for which side. Therefore we use `on` keyword. If no `on` keyword is
used, the compiler will assume that the file is for both sides (shared).

```ggw
on frontend;
```

```ggw
on backend;
```

```ggw
on shared;
```

The `on` keyword can be used with module and class declarations in one statement.

```ggw
module math on shared;
```

```ggw
class Point on shared {
    x: int;
    y: int;
}
```

The `on` keyword can be used with .mod files to specify for the whole directory.

```ggw
module math on shared;
```

Now the fun part, the wiring. The wiring is done with the `wire` keyword. The `wire` keyword can be used on classes,
instances, methods and properties/varaibles. When using the `wire` keyword, an adapter must be specified otherwise
the compiler will use `HTTP` as default.

```ggw
on backend;

func add(a: int, b: int): int wire WebSockets {
    return a + b;
}
```

```ggw
on frontend;

let result = wire.add(1, 2);
```

Obviously without some kind of authentication this is a security risk. So the language will have a way to handle
authentication and authorization.

First define a way to authenticate. Therefore we use another keyword called `grant` which grants the current connection
authentication. using `with` behind the `grant` keyword specifies roles which are allowed to access the method.
With can be used with arrays of strings.
```ggw
on backend

func login(username: string, password: string): bool wire WebSockets {
    if username == "admin" && password == "admin" {
        grant with ["admin"];
    } else if username == "user" && password == "user" {
        grant with ["user"];
    } else {
        // granting nothing will return an login error
        return false;
    }
    
    return true;
}
```

```ggw
on backend;

func add(a: int, b: int): int wire WebSockets authenticated("admin") {
    return a + b;
}
```

```ggw
on frontend;

let result = wire.add(1, 2); // this will return an error because the user is not authenticated

let success = wire.login("admin", "admin"); // this will authenticate the user as admin

let result = wire.add(1, 2); // this will return 3
```

Another cool thing: as said, the `wire` keyword can work with variables and values, too and even with classes. The
de/serialization is done automatically. By default using JSON. But you can specify the format (maybe later).

```ggw
on backend;

let x = 42 wire WebSockets;

class Point {
    x: int;
    y: int;
}

let point = new Point(1, 2) wire WebSockets;
```

```ggw
on frontend;

let x = wire.x;
let point = wire.point;
```

Wiring classes with functions work to a certain extent. The compiler will automatically create a method for each function
as long as the class is shared. If the class is not shared, the compiler will throw an error.

```ggw
on shared;

class Math {
    func add(a: int, b: int): int {
        return a + b;
    }
}
```

```ggw
on backend;

let math = new Math() wire WebSockets;
```

```ggw
on frontend;

let result = wire.math.add(1, 2);
```