package staticUtil

import (
	"fmt"
	"time"
)

// Transforms a date to filepath
func GenerateDatePath() string {
	now := time.Now()
	return fmt.Sprintf("%d/%d/%d/", now.Year(), now.Month(), now.Day())
}

// Generate date and time string
func GetDate() string {
	n := time.Now()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second())
}
