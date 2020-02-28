
# go-float-decimal
Example usage of float and decimal in Golang

## Examples

```shell
>>> 165 * 1.40
230.99999999999997
```

```shell
>>> 1.01 - 0.99
0.020000000000000018
```

```shell
>>> 1e16 * (1.01 - 0.99)
200000000000000
```

```shell
>>> 1e16 * (1.01 - 0.99) - 0.01
199999999999999.99
```

---

## Links

- https://husobee.github.io/money/float/2016/09/23/never-use-floats-for-currency.html
- https://stackoverflow.com/questions/3730019/why-not-use-double-or-float-to-represent-currency#3730040
- https://godoc.org/github.com/shopspring/decimal
