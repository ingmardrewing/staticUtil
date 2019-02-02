package staticUtil

import (
	"fmt"
	"time"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

var (
	timeNow = time.Now
)

// Transforms a date to filepath
func GenerateDatePath() string {
	n := timeNow()
	return fmt.Sprintf("%d/%d/%d/", n.Year(), n.Month(), n.Day())
}

// Generate date and time string
func GetDate() string {
	n := timeNow()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second())
}

func DirExists(path string) bool {
	exits, _ := fs.PathExists(path)
	return len(path) > 0 && exits
}

func MakeSrcSet(p staticIntf.Page) string {
	return fmt.Sprintf("%s 390w, %s 800w",
		p.ThumbnailUrl(),
		p.ImageUrl())
}
