# Fython Project
Fython is a functional flavour of Python.


# Fython vs Python

To achieve the functional twist some changes need to occur,
while we maintain the flavour

## Data Types

| Fython   | Python Equivalent |
| -------- | ----------------- |
| bool     | bool              |
| int      | int               |
| float    | float             |
| str      | str (immutable)   |
| bytes    | bytes             |
| tuple    | tuple             |
| set      | frozenset         |
| dict     | frozendict        |
|          | list              |
| Optional | None              |


### Optional & Error Handling

Optional is an enum wrapper with members `PRESENT(x)` and `NONE`.

# Multiline Lambdas
```python
func = lambda(x, y, z):
    return x + y + z
```
