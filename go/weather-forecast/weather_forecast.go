// Package weather provides tools to
// see city's condition.
package weather

// CurrentCondition provides actual condition.
var CurrentCondition string
// CurrentLocation provides actual location.
var CurrentLocation string

// Forecast returns a string with current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
