package sorting

import (
  "fmt"
  "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
  return "This is the number " + strconv.FormatFloat(f, 'f', 1, 64)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
  return "This is a box containing the number " + strconv.FormatFloat(float64(nb.Number()), 'f', 1, 64)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
  if _, ok := fnb.(FancyNumber); !ok {
    return 0
  }
  value, err := strconv.Atoi(fnb.Value())
  if err != nil {
    return 0
  }
  return value
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
  return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
  switch v := i.(type) {
  case int:
    return DescribeNumber(float64(v))
  case float64:
    return DescribeNumber(v)
  case NumberBox:
    return DescribeNumberBox(v)
  case FancyNumberBox:
    return DescribeFancyNumberBox(v)
  default:
    return "Return to sender"
  }
}
