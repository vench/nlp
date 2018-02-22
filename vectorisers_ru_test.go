package nlp


import (
	"testing"
	"fmt"
)

func TestModeWord(t *testing.T) {

	fmt.Println("TestModeWord")

	tests := []struct{
		test string
		eval string} {

		{eval: "московского", test :"московский"},
		{eval: "московскому", test: "московский"},
		{eval: "московском", test: "московский"},
		{eval: "Рокоссовскому", test: "Рокоссовский"},
		{eval: "авиаудара", test: "авиаудар"},

		{eval: "авиаудара", test: "авиаудар"},
		{eval: "авиаударам", test: "авиаудар"},
		{eval: "авиаударами", test: "авиаудар"},
		{eval: "авиаударах", test: "авиаудар"},
		{eval: "авиаударе", test: "авиаудар"},
		{eval: "авиаударов", test: "авиаудар"},
		{eval: "авиаудары", test: "авиаудар"},

		{eval: "путина", test: "путин"},
		{eval: "авиатора", test: "авиатор"},
		{eval: "авиатору", test: "авиатор"},
		{eval: "авиаторы", test: "авиатор"},
		{eval: "авиатором", test: "авиатор"},
		//{eval: "тора", test: "тора"},

		{eval: "колокольчику", test: "колокольчик"},
		{eval: "колокольчиков", test: "колокольчик"},
		{eval: "колокольчика", test: "колокольчик"},
		{eval: "колокольчиками", test: "колокольчик"},
		{eval: "колокольчиком", test: "колокольчик"},
		{eval: "колокольчиками", test: "колокольчик"},
		{eval: "паровозика", test: "паровозик"},
		//{eval: "чика", test: "чика"},


		{eval: "паромщика", test: "паромщик"},
		{eval: "паромщику", test: "паромщик"},
		{eval: "паромщикам", test: "паромщик"},
		{eval: "паромщиков", test: "паромщик"},
		{eval: "паромщиком", test: "паромщик"},
	}

	for _, ts := range tests {

		res := MorfModeWord(ts.eval)
		//fmt.Println(res)
		if res != ts.test {
			t.Fatalf("ModeWord(%s) != %s (r: %s)\n", ts.eval, ts.test, res )
		}
	}
}
