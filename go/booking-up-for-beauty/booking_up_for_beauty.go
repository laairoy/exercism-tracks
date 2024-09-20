package booking

import (
	"time"
)

func checkDate(layout, date string) time.Time {
  result,err := time.Parse(layout, date)
  if err != nil {
    panic(err)
  }
  return result
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
  return checkDate("1/2/2006 15:04:05", date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
  result := checkDate("January 2, 2006 15:04:05", date)

  if time.Now().Compare(result) == -1 {
    return false
  }
  return true
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
  result := checkDate("Monday, January 2, 2006 15:04:05", date) 
  
  if result.Hour() >= 12 && result.Hour() <= 18 {
    return true
  }

  return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
  result := Schedule(date)
  return "You have an appointment on " + result.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
  return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
