# Algorithms and Data Structures 2 - Practical Works

This repository contains practical assignments from the **Algorithms and Data Structures 2** course at FIUBA (Facultad de Ingeniería, Universidad de Buenos Aires). It includes implementations of abstract data types (TDAs) in Go and practical problems in Go and Python.

## Repository Structure

- **`tdas/`**: Implementations of abstract data types in Go.
  - **`cola_prioridad/`**: Priority queue with heap implementation.
  - **`diccionario/`**: Dictionary implementations (hash table, ordered dictionary with ABB).
  - **`lista/`**: Linked list implementation.
  - **`pila/`**: Stack implementations (dynamic).
  - Each module includes tests (`*_test.go`).
- **`tp2/`**: Practical Work 2 - File reader and data structures usage.
  - **`lector/`**: Reader for datasets and data structures.
- **`tp3/`**: Practical Work 3 - Monopoly-like game simulation in Python.
  - `algopoli.py`: Main game logic.
  - `grafo.py`: Graph implementation for board.
  - `biblioteca.py`: Utility functions.
  - `pila.py`: Stack implementation.
  - `entrega.mk`: Makefile for delivery.
- **`tp3_others/`**: Additional files for TP3, including scripts and data files.

## Prerequisites

- **Go**: Version 1.19 or higher.
- **Python**: Version 3.8 or higher.

## Installation and Execution

### TDAs (Go)

1. Navigate to the `tdas` directory:
   ```bash
   cd tdas
   ```

2. Run tests for a specific TDA:
   ```bash
   go test ./cola_prioridad
   go test ./diccionario
   go test ./lista
   go test ./pila
   ```

3. Run all tests:
   ```bash
   go test ./...
   ```

### TP2 (Go)

1. Navigate to `tp2`:
   ```bash
   cd tp2
   ```

2. Run the main program:
   ```bash
   go run main.go
   ```

### TP3 (Python)

1. Navigate to `tp3`:
   ```bash
   cd tp3
   ```

2. Run the game:
   ```bash
   python algopoli.py
   ```

3. Use the Makefile for specific tasks:
   ```bash
   make entrega
   ```

## Topics Covered

- Abstract data types: stacks, queues, lists, dictionaries.
- Data structures: heaps, hash tables, binary search trees.
- Algorithms: sorting, searching, graph traversal.
- File I/O and data processing.
- Game simulation and object-oriented programming.

## Author

- **Enrique Heller** - [EnriHeller](https://github.com/EnriHeller)

## License

This repository is for educational purposes. Use and modify freely, but cite the source for academic work.

## Contributing

Suggestions and improvements are welcome via issues or pull requests.