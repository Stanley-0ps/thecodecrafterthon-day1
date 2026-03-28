# 🧮 CLI Calculator (Go)

A simple command-line calculator built with Go.
It performs basic arithmetic operations and handles invalid user input gracefully.

---

## ✨ Features

* ➕ Addition
* ➖ Subtraction 
* ✖️ Multiplication
* ➗ Division with zero-check
* 🚫 Input validation (prevents crashes on invalid input)
* 💻 Interactive CLI menu

---

## 🚀 Getting Started

### Prerequisites

* Install Go (version 1.18 or later recommended)

### Run the program

```bash
go run main.go
```

---

## 📌 Usage

1. Run the program
2. Select an operation from the menu
3. Enter two numbers when prompted

Example:

```
Select operation:
1. Addition
2. Division
Enter choice: 1

Input first number: 5
Input second number: 3

Result: 8
```

---

## ⚠️ Error Handling

This calculator is designed to handle common user mistakes:

### Invalid number input

```
Input first number: cat
Error: please enter a valid number
```

### Division by zero

```
Error: division by zero
```

The program avoids crashes and provides clear feedback instead.

---

## 🧠 Project Structure

```
thecodecrafterthon-day1/
├── main.go          # Entry point and CLI logic
├── README.md        # Programme Description
├── communication.go    # presentation of options  
├── math.go        # Arithmetic functions (Addition, Division, etc.)
```

---

## 📖 Example Functions

```go
func Addition(a, b float64) float64 {
    return a + b
}

func Division(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

---

## 🎯 Future Improvements

* Add more operations (modulus, power)
* Support for command-line arguments
* Unit tests

---

## 📄 License

This project is open-source and free to use.
