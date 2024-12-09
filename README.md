# Multi-Language FizzBuzz Implementation

This repository demonstrates FizzBuzz solutions implemented in four programming languages (JavaScript, Python, Rust, and Go), using both iterative and recursive approaches.

## What is FizzBuzz?

FizzBuzz is a classic programming problem:

- Print numbers from 0 to 100
- For multiples of 3, print "Fizz!" instead of the number
- For multiples of 5, print "Buzz!" instead of the number
- For multiples of both 3 and 5, print "FizzBuzz!"

## Implementations

### JavaScript

Located in `js/main.js`:

- Iterative solution using a for loop
- Recursive solution using a helper function

Example usage:

```javascript
// Run either solution:
const array = Array.from(Array(101).keys());
recurse(array, 0); // For recursive solution
```

<code_block_to_apply_changes_from>

```python
array = list(range(101))
recurse(array, 0)  # For recursive solution
```

### Rust

Located in `rust/src/main.rs`:

- Vector-based iterative solution
- Memory-safe recursive implementation

Example usage:

```rust
fn main() {
    let array: Vec<i32> = (0..101).collect();
    recurse(&array, 0);  // For recursive solution
}
```

### Go

Located in `golang/main.go`:

- Slice-based iterative solution
- Clean recursive implementation with error handling

Example usage:

```go
func main() {
    array := make([]int, 101)
    for i := range array {
        array[i] = i
    }
    recurse(array, 0)  // For recursive solution
}
```

### Python

Located in `python/main.py`:

- Clean iterative solution using a for loop
- Recursive implementation with base case handling

Example usage:

```python
array = list(range(101))
recurse(array, 0)  # For recursive solution
```

## Project Structure

```
.
├── js/
│   ├── main.js
│   └── package.json
├── python/
│   └── main.py
├── rust/
│   ├── src/
│   │   └── main.rs
│   └── Cargo.toml
└── golang/
    ├── main.go
    └── go.mod
```

## Key Differences Between Implementations

1. **JavaScript**

   - Uses Array.from() for array creation
   - Demonstrates both functional and imperative approaches

2. **Python**

   - Utilizes list comprehension for array creation
   - Features cleaner conditional logic with elif statements

3. **Rust**

   - Implements memory-safe array handling
   - Uses Vec<i32> for type-safe integer collection

4. **Go**
   - Employs slice-based array manipulation
   - Features explicit error handling patterns

## Running the Examples

### JavaScript

```bash
node js/main.js
```

### Python

```bash
python python/main.py
```

### Rust

```bash
cd rust
cargo run
```

### Go

```bash
cd golang
go run main.go
```

## License

ISC License - See package.json for details

```

This README provides a comprehensive overview of your FizzBuzz implementations, including:
- Detailed problem description
- Implementation details for each language
- Project structure
- Key differences between implementations
- Running instructions
- Code examples for each language
```
