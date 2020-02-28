package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	floatCalc()
	fmt.Println("")
	floatPrint()
	fmt.Println("")
	fmt.Println("")

	decimalCalc()
	fmt.Println("")
	decimalPrint()
	fmt.Println("")
	furtherDecimalCalc()
	fmt.Println("")

	// DivisionPrecision is the number of decimal places in the result when it doesn't divide exactly.
	// [default] decimal.DivisionPrecision = 16
	reducedDecimalPrecision()
	fmt.Println("")
	fmt.Println("")
}

func floatCalc() {
	fmt.Println("*** Float calculations ***")

	var x float64 = 165.0
	var y float64 = 1.4
	result := x * y
	fmt.Println("expected \t 165.0 * 1.4 = 231")
	fmt.Printf("actual \t\t %.16f * %.16f = %.16f\n", x, y, result)

	x = 1.01
	y = 0.99
	result = x - y
	fmt.Println("expected \t 1.01 - 0.99 = 0.020000000000000018")
	fmt.Printf("actual \t\t %.16f * %.16f = %.24f\n", x, y, result)

	var a float64 = 1e16
	var b float64 = 1.01
	var c float64 = 0.99
	var d float64 = 0.01
	result = a * (b - c)
	fmt.Println("expected \t 1e16 * (1.01 - 0.99) = 200000000000000")
	fmt.Printf("actual \t\t %.16f * (%.16f - %.16f) = %.16f\n", a, b, c, result)
	result = a*(b-c) - d
	fmt.Println("expected \t 1e16 * (1.01 - 0.99) - 0.01 = 199999999999999.99")
	fmt.Printf("actual \t\t %.16f * (%.16f - %.16f) - %.16f = %.16f\n", a, b, c, d, result)
}

func floatPrint() {
	fmt.Println("*** Printing floats ***")
	f := 0.020000000000000018
	fmt.Println("expected \t 0.020000000000000018")
	fmt.Printf("actual \t\t %.24f\n", f)
}

func decimalCalc() {
	fmt.Println("*** Decimal calculations ***")
	x := decimal.NewFromFloat(165.0)
	y := decimal.NewFromFloat(1.4)
	result := x.Mul(y)
	fmt.Println("expected \t 165.0 * 1.4 = 231")
	fmt.Printf("actual \t\t %s * %s = %s\n", x.String(), y.String(), result.String())

	x = decimal.NewFromFloat(1.01)
	y = decimal.NewFromFloat(0.99)
	result = x.Sub(y)
	fmt.Println("expected \t 1.01 - 0.99 = 0.02")
	fmt.Printf("actual \t\t %s * %s = %s\n", x.String(), y.String(), result.String())

	a := decimal.NewFromFloat(1e16)
	b := decimal.NewFromFloat(1.01)
	c := decimal.NewFromFloat(0.99)
	d := decimal.NewFromFloat(0.01)
	result = a.Mul(b.Sub(c))
	fmt.Println("expected \t 1e16 * (1.01 - 0.99) = 200000000000000")
	fmt.Printf("actual \t\t %s * (%s - %s) = %s\n",
		a.String(), b.String(), c.String(), result.String())

	result = a.Mul(b.Sub(c)).Sub(d)
	fmt.Println("expected \t 1e16 * (1.01 - 0.99) - 0.01 = 199999999999999.99")
	fmt.Printf("actual \t\t %s * (%s - %s) - %s = %s\n",
		a.String(), b.String(), c.String(), d.String(), result.String())
}

func decimalPrint() {
	fmt.Println("*** Printing decimals ***")
	f := 0.020000000000000018
	d := decimal.NewFromFloat(f)
	fmt.Println("expected \t 0.020000000000000018")
	fmt.Printf("actual \t\t %s\n", d.String())
}

func furtherDecimalCalc() {
	fmt.Println("*** Further decimal calculations ***")

	x := decimal.NewFromFloat(2.0)
	y := decimal.NewFromFloat(3.0)
	result := x.Div(y)
	fmt.Println("expected \t 2 / 3 = 0.6666666666666667")
	fmt.Printf("actual \t\t %s / %s = %s\n", x.String(), y.String(), result.String())

	x = decimal.NewFromFloat(2.0)
	y = decimal.NewFromFloat(30000.0)
	result = x.Div(y)
	fmt.Println("expected \t 2 / 30000 = 0.0000666666666667")
	fmt.Printf("actual \t\t %s / %s = %s\n", x.String(), y.String(), result.String())

	x = decimal.NewFromFloat(20000.0)
	y = decimal.NewFromFloat(3.0)
	result = x.Div(y)
	fmt.Println("expected \t 20000 / 3 = 61666.6666666666666667")
	fmt.Printf("actual \t\t %s / %s = %s\n", x.String(), y.String(), result.String())
}

func reducedDecimalPrecision() {
	fmt.Println("*** Reduced decimal precision from 16 to 8 ***")

	decimal.DivisionPrecision = 8
	x := 20000.0
	y := 3.0
	result := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y))
	fmt.Println("expected \t 20000 / 3 = 61666.66666667")
	fmt.Printf("actual \t\t %.16f / %.16f = %s\n", x, y, result.String())
}
