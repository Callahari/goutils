# Random Utilities

## Overview

This package provides a collection of utility functions to generate random values, including integers, float32, float64, and dates within a specified range.

## Functions

### RandomIntBetween

Generates a random integer between a specified minimum and maximum range.

#### Function Signature

```go
func RandomIntBetween(min, max int) int
```

#### Parameters

- `min` (int): The minimum value of the range (inclusive).
- `max` (int): The maximum value of the range (inclusive).

#### Returns

- `int`: A random integer between `min` and `max`.

### RandomFloat32Between

Generates a random float32 number between a specified minimum and maximum range.

#### Function Signature

```go
func RandomFloat32Between(min, max float32) float32
```

#### Parameters

- `min` (float32): The minimum value of the range.
- `max` (float32): The maximum value of the range.

#### Returns

- `float32`: A random float32 number between `min` and `max`.

### RandomFloat64Between

Generates a random float64 number between a specified minimum and maximum range.

#### Function Signature

```go
func RandomFloat64Between(min, max float64) float64
```

#### Parameters

- `min` (float64): The minimum value of the range.
- `max` (float64): The maximum value of the range.

#### Returns

- `float64`: A random float64 number between `min` and `max`.

### RandomDateBetween

Generates a random date between a specified start date and end date.

#### Function Signature

```go
func RandomDateBetween(startDate, endDate time.Time) time.Time
```

#### Parameters

- `startDate` (time.Time): The start date of the range.
- `endDate` (time.Time): The end date of the range.

#### Returns

- `time.Time`: A random date between `startDate` and `endDate`.

## Example Usage

Here is an example of how to use these functions in your Go code:

```go
package main

import (
	"fmt"
	"github.com/Callahari/goutils/goutils"
	"time"
)

func main() {
	// Random Int
	fmt.Println("Random Int:", goutils.RandomDateBetween(1, 10))

	// Random Float32
	fmt.Println("Random Float32:", goutils.RandomFloat32Between(1.0, 10.0))

	// Random Float64
	fmt.Println("Random Float64:", goutils.RandomFloat64Between(1.0, 10.0))

	// Random Date
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	fmt.Println("Random Date:", goutils.RandomDateBetween(startDate, endDate))
}

```

## Notes

- Ensure that you seed the random number generator using `rand.Seed` with a value like `time.Now().UnixNano()` for unique random numbers in each run.
- The functions will generate values in the inclusive range from `min` to `max`.

## Contact

For any questions or feedback, use github functionality