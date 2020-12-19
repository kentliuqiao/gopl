package main

import (
	"flag"
	"fmt"
)

type Celcius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32.0) * 5.0 / 9.0)
}

func KToC(k Kelvin) Celcius {
	return Celcius(k - 273.15)
}

func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c) // %g	%e for large exponents, %f otherwise. Precision is discussed below. %e	scientific notation, e.g. -1.234456e+78
}

type celciusFlag struct {
	Celcius
}

func (f *celciusFlag) Set(s string) error {
	var unit string
	var value float64

	// no need to handle error here, since err will not fit any case below
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celcius = Celcius(value)
		return nil
	case "F", "°F":
		f.Celcius = FToC(Fahrenheit(value))
		return nil
	case "K", "ºK":
		f.Celcius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelciusFlag(name string, value Celcius, usage string) *Celcius {
	f := celciusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celcius
}
