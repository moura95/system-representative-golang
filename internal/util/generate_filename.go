package util

import (
	"path/filepath"
	"strconv"
	"time"
)

func GenerateFilename(representativeID int32, filename, dir string) string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	extension := filepath.Ext(filename)
	return strconv.Itoa(int(representativeID)) + "/" + dir + "/" + strconv.FormatInt(timestamp, 10) + extension
}
