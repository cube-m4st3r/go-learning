# go-learning

Small collection of simple Go example programs used for learning and experimentation.

## Projects

- `bank/` - Small banking example that demonstrates file operations and basic I/O. Contains `balance.txt` and a `fileops` helper package.
- `investment_calculator/` - Example showing basic investment calculations.
- `profit_calculator/` - Simple profit calculation example.

## Prerequisites

- Go (the dev container includes a recent Go toolchain).

## Running

From the repository root you can run any example with `go run`:
```bash
# run the bank example
go run ./bank

# run the investment calculator
go run ./investment_calculator

# run the profit calculator
go run ./profit_calculator
```

## Notes

- Each example is a small self-contained module. Inspect the `go.mod` files in the subfolders for module-specific settings.
- Feel free to modify and experiment — these programs are intended for learning.
