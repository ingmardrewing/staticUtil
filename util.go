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

func MakeMicroSrcSet(p staticIntf.Page) string {
	if len(p.ThumbnailUrl()) <= 5 {
		return ""
	}
	return fmt.Sprintf("%s 2x", p.ThumbnailUrl())
}

func MakeSrcSet(p staticIntf.Page) string {
	if len(p.ImageUrl()) <= 5 ||
		p.ThumbnailUrl() == p.ImageUrl() {
		return ""
	}
	return fmt.Sprintf("%s 2x", p.ImageUrl())
}
