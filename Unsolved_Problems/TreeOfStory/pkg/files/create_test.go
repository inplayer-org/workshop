package files

import (
	"testing"
)


type Tests struct{
	ReturnedError string
	ExpectedError string
}


func TestPath(t *testing.T) {
	path,err:=Path("testingfolder")

	if err==nil {
		if path != "/home/gligor/code/go/src/repo.inplayer.com/workshop/Unsolved_Problems/TreeOfStory/testingfolder" {

		t.Error("Expected home/gligor/code/go/src/repo.inplayer.com/workshop/Unsolved_Problems/TreeOfStory/testingfolder, got ", path)
		}
	}

	}

func TestCreateDirectory(t *testing.T) {

	testPaths:=make(map[string]string)

	testPaths["test"]=""
	testPaths["test/teeest"]=""
	testPaths["../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../sfaa2"]="Can't create directory"


for k,v:=range testPaths {
	path,_:=Path(k)

	err:=CreateDirectory(path)

	if err!=nil {
		if err.Error() != v {
			t.Error("Expected ", v, ",got", err.Error())
		}
	}else{
		if ""!=v{
			t.Error("Expected ", v, ",got", "")
		}
	}
}
}

