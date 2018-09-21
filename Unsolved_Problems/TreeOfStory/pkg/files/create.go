package files

import (
	"os"
)


//creating directory in given path(string) if not exist
func CreateDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}





