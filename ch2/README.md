# Predeclared Types and Declarations

**_The zero value_**

Go assigns a default _zero value_ to any variable that is declared but not assigned a value.

**_Booleans_**

The bool type represents Boolean variables. bool type can have one of two values: true or false. The zero value of a bool is false

**_Integers_**

The zero value for all of the integer types is 0.

| Type name | Value range                                 |
| --------- | ------------------------------------------- |
| int8      | -127 to 127                                 |
| int16     | -32768 to 32767                             |
| int32     | -2147483648 to 2147483647                   |
| int64     | -9223372036854775808 to 9223372036854775807 |
| uint8     | 0 to 255                                    |
| uint16    | 0 to 65536                                  |
| uint32    | 0 to 4294967295                             |
| uint64    | 0 to 18446744073709551615                   |

**The special integer types**

Go does have some special names for integer types. A byte is an alias for uint8; it is legal to assign, compare or perform mathematical operations between byte and uint8. However, you rarely see uint8 used in Go code; just call it byte.

The second special name is int. On a 32-bit CPU, int is a 32-bit signed integer like int32. On most 64-bit CPU, int is a 64-bit signed integer, just like an int64.

The third special name is uint. It follows the same rules as int, only it is unsigned (the value are always 0 or positive)

There are two other special names for integer types, rune and uintptr.

**Choosing which integer to use**

- If you are working with a binary file format or network protocol that has an integer of specific size or sign, use the corresponding integer type.
- If you are writing a library function that should work with any integer type, take advantage of Go's generics support and use a generic type parameter to represennt any integer type
- In all other cases, just use int.

**Integer operators**

Go integers support the usual arithmetic operators: +, -, \*, /, %

You compare integers with ==, !=, >, >=, <, <=.

Go also has bit-manipulation operators for integers. You can bit shift left and right with << and >>, or do bit masks with & (logical AND), | (logical OR), ^ (logical XOR), and &^ (logical AND NOT). You can also combine all of the logical operators with = to modify a variable: &=, |=, ^=, &^=, <<=, >>=.

**Floating point types**

Like integer types, the zero value for the floating point type is 0.

| Type name | Largest absolute value                         | Smallest(nonzero) absolute value               |
| --------- | ---------------------------------------------- | ---------------------------------------------- |
| float32   | 3.40282346638528859811704183484516925440e+38   | 1.401298464324817070923729583289916131280e-45  |
| float64   | 1.797693134862315708145274237317043567981e+308 | 4.940656458412465441765687928682213723651e-324 |

**_A taste of strings and runes_**

Go includes strings as a built-in type. The zero value of a string is the empty string. Go support Unicode; strings are compared for equality using ==, difference !=, or ordering with >, >=, < or <=. They are concatenated by using the + operator

Strings in Go are immutable. You can reassign the value of a string variable, but you cannot change the value of a string that is assigned to it.

Go also has a type that represents a single code point. The _rune_ type is an alias for int32 type.

> If you are referring to a character, use the rune type, not the int32 type. They might be the same to the compiler, but you want to use the type that clarifies the intent of your code.

**_var versus :=_**

- When initializing a variable to its zero value, use var x int . This makes it clear that zero value is intended.
- When assigning an untyped constant or a literal to a variable and the default type for the constant or literal isn't the type you want for the variable, use the long var form with the type specified.\
  While it is legal to use a type conversion to specify the type of the value and use := to write x := byte(20), it is idiomatic to write var x byte = 20.
- Because := allows you to assign to both new and existing variable, it sometimes creates new variable when you think you are reusing existing one. In those situations, explicitly declare all \
  of your new variables with var to make it clear which variables are new, and then use the assignment operator (=) to assign values to both new and old variables.

**_Using const_**

_const_ is a way to declare that a value is immutable. However _const_ in Go is very limited. Constants in Go are a way to give names to literals. \
They can only hold values that the compiler can figure out at compile time.

There is no way in Go to declare that a variable is immutable.

This means they can be assigned:

- Numeric literals
- true and false
- strings
- runes
- The built-in functions complex, real, imag, len, and cap
- Expression that consist of operators and the preceding values

**_Typed and Untyped Constants_**

Constants can be typed or untyped. An untyped constant works exactly like a literal; it has no type of its own but does have a default type that is used when no other type can be inferred. \
A typed constant can be directly assigned only to a variable of that type

**_Naming Variables and Constnats_**

- Idiomatic Go uses camel case when an identifier name consists of multiple words.
- In many languages, constant are always written in all uppercase letters separated by underscores. Go does not follow this pattern. Go uses tha case of the first letter in the name of a package-level declaration to determine if the item is accessible outside the package.
- Within a function, favor short variable names. The smaller the scope for a variable, the shorter the name that's used for it. It is common in go to see single-letter variable names used with for loops.
