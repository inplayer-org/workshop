package parser_test

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"testing"
)

type Tests struct{
	Argument string
	ExpectedResult Results
}

type Results struct {
	Str string
	Err string
}

func TestFromAnyToHashTag(t *testing.T){
	var TestTable = []Tests{
		{Argument:"",ExpectedResult:Results{"","tried to convert empty string"}},
		{Argument:"%25%%#&%25##",ExpectedResult:Results{"","string is only consisted of #, %25,% and &"}},
		{Argument:"test",ExpectedResult:Results{"#test",""}},
		{Argument:"#test",ExpectedResult:Results{"#test",""}},
		{Argument:"##test",ExpectedResult:Results{"#test",""}},
		{Argument:"%25test",ExpectedResult:Results{"#test",""}},
		{Argument:"%%25test",ExpectedResult:Results{"#test",""}},
		{Argument:"%25%test",ExpectedResult:Results{"#test",""}},
		{Argument:"&test",ExpectedResult:Results{"#test",""}},
		{Argument:"&&test",ExpectedResult:Results{"#test",""}},
		{Argument:"&&&test",ExpectedResult:Results{"#test",""}},
		{Argument:"%25%#&test",ExpectedResult:Results{"#test",""}},
		{Argument:"&#%25test",ExpectedResult:Results{"#test",""}},
	}

	for _,test := range TestTable{
		str,err := parser.FromAnyToHashTag(test.Argument)
		if err==nil{
			if str!=test.ExpectedResult.Str || ""!=test.ExpectedResult.Err{
				t.Errorf("Failed test for string:%s,\n returned string: %s \t error: %s\n expected string: %s \t error: %s",
					test.Argument,
					str,"",
					test.ExpectedResult.Str,test.ExpectedResult.Err)
			}
		}else {
			if str != test.ExpectedResult.Str || err.Error() != test.ExpectedResult.Err {
				t.Errorf("Failed test for string:%s,\n returned string: %s \t error: %s\n expected string: %s \t error: %s",
					test.Argument,
					str, err,
					test.ExpectedResult.Str, test.ExpectedResult.Err)
			}
		}
	}

}

