package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
  return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`).MatchString(text)
}

func SplitLogLine(text string) []string {
  return regexp.MustCompile(`<([-~\*=]*)>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
  var count int
  re := regexp.MustCompile(`".*(?i)(password).*"`)
  for _, line := range lines {
    if re.MatchString(line) {
      count++
    }
  }
  return count
}

func RemoveEndOfLineText(text string) string {
  return regexp.MustCompile(`(end-of-line\d*)`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
  re := regexp.MustCompile(`User\s+(\w+)`)
  for index, line := range lines {
    groups := re.FindStringSubmatch(line)
    if groups != nil {
      lines[index] = fmt.Sprintf("[USR] %s %s", groups[1], line)
    }
  }
  return lines
}
