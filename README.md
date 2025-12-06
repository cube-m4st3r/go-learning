# go-learning

Small collection of simple Go example programs used for learning and experimentation.

## Projects

- `bank/` - Small banking example that demonstrates file operations and basic I/O. Contains `balance.txt` and a `fileops` helper package.
- `investment_calculator/` - Example showing basic investment calculations.
- `profit_calculator/` - Simple profit calculation example.

## Prerequisites

- Go (the dev container includes a recent Go toolchain).

> [!IMPORTANT]  
> The current configuration for the devcontainer does **NOT** support out of the box windows runtime. You might have do adjust it for your specific needs.

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

## Further plans for this repo

- Implement .git hooks or use the pre-commit framework
- Update the devcontainer with a custom Dockerfile to build a container with go and pre-commit framework installed (if so chosen)
- Create better GitHub Actions aka. CI/CD
- Write documentation and explanations for each section and mini project

## License

This project is provided under the MIT License — see `LICENSE` for full terms.
