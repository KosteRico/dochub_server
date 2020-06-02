package indexing

import "golang.org/x/text/encoding/charmap"

func decodeWin1251(b []byte) string {
	bytes, _ := charmap.Windows1251.NewDecoder().Bytes(b)
	return string(bytes)
}
