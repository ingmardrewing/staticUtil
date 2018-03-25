package staticUtil

import (
	"fmt"
	"time"

	"github.com/ingmardrewing/fs"
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

func DirExists(path string) bool {
	exits, err := fs.PathExists(path)
	if err != nil {
		return false
	}
	return len(path) > 0 && exits
}
