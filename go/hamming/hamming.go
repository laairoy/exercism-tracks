package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
  if len(a) == len(b) {
    var count int
    for index := range a {
      if a[index] != b[index] {
        count++
      }
    }
    return count, nil
  }

  return 0, errors.New("Invalid input!")
}
