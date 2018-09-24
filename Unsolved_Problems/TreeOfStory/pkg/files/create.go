package files

import (
	"os"
	"fmt"
	"io"
	"os/user"
	"repo.inplayer.com/workshop/Unsolved_Problems/TreeOfStory/pkg/errors"
)

var path = "/home/darko/go/src/repo.inplayer.com/workshop/Unsolved_Problems/TreeOfStory/user"


//creating directory in given path(string) if not exist
func CreateDirectory(path string) error{
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return errors.NewMyError("Can't create directory")
		}
	}
	return nil
}

func CreateFile(path string) error {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) { return errors.NewMyError("Can't create file") }
		defer file.Close()
	}


	fmt.Println("==> done creating file", path)
return nil
}


func ReadFile(path string) error{
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return errors.NewMyError("Can't open file")}
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			errors.NewMyError("Can't read file")
		}
	}

	fmt.Println("==> done reading from file")
	fmt.Println(string(text))
	return nil
}

func DeleteFile(path string) error{
	// delete file
	var err = os.Remove(path)
	if isError(err) { return errors.NewMyError("Can't delete file") }

	fmt.Println("==> done deleting file")
	return nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func Path(inFile string)(string,error){
	user,err:=user.Current()

	if err!=nil {
		return "",errors.NewMyError("Can't find HomeDir")
	}

	return user.HomeDir+"/code/go/src/repo.inplayer.com/workshop/Unsolved_Problems/TreeOfStory/"+inFile,nil
}