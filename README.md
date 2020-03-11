
# go-float-decimal
Example usage of float and decimal in Golang

## Examples

```shell
165 * 1.40
float64 >>> 230.9999999999999716
decimal >>> 231
```

```shell
1.01 - 0.99
float64 >>> 0.020000000000000017763568
decimal >>> 0.020000000000000018
```

```shell
>>> 1e16 * (1.01 - 0.99)
float64 >>> 200000000000000.1875000000000000
decimal >>> 200000000000000
```

```shell
>>> 1e16 * (1.01 - 0.99) - 0.01
float64 >>> 200000000000000.1875000000000000
decimal >>> 199999999999999.99
```

---

## Links

- https://husobee.github.io/money/float/2016/09/23/never-use-floats-for-currency.html
- https://stackoverflow.com/questions/3730019/why-not-use-double-or-float-to-represent-currency#3730040
- https://godoc.org/github.com/shopspring/decimal
