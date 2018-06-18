package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

//valid actions: verify, shuffle,
type MyEvent struct {
	StringA string `json:"stringA,omitempty"`
	StringB string `json:"stringB,omitempty"`
	Shuffle string `json:"shuffle,omitempty"`
	Action  string `json:"action",omitempty`
	Debug   bool   `json:"debugon,omitempty"`
}

type MyResponse struct {
	Result      map[string]interface{} `json:"result"`
	DebugOutput []string               `json:"debugOutput,omitempty"`
}

var debugOn bool = false
var debugOut []string

func main() {
	lambda.Start(HandleLambdaEvent)
}

func (m MyEvent) string() string {
	return fmt.Sprintf(m.Action, m.StringA, m.StringB, m.Shuffle)
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	// str1 := "abc"
	// str2 := "def"
	// strValid := "abdecf"
	// strNotValid := "acbfed"
	debugOn = event.Debug
	debugOut = []string{}
	var resp MyResponse

	//check debug
	if debugOn {
		debugOut = []string{"INFO: Start debug", event.string()}
		resp = MyResponse{
			Result:      make(map[string]interface{}),
			DebugOutput: debugOut,
		}
	} else {
		resp = MyResponse{
			Result: make(map[string]interface{}),
		}
	}

	switch event.Action {
	case "shuffle":
		{
			resp.Result["shuffle"] = event.StringA + event.StringB
		}
	case "verify":
		{
			if shuffle(event.StringA, event.StringB, event.Shuffle) {
				resp.Result["Valid"] = true
			} else {
				resp.Result["Valid"] = false
			}
		}
	default:
		{
			resp.Result["error"] = "Your Action (" + event.Action + ") was not valid, please choose verify or shuffle."
		}
	}

	if debugOn {
		resp.DebugOutput = debugOut
	}

	return resp, nil

}

func shuffle(str1, str2, str3 string) bool {
	if debugOn {
		debugOut = append(debugOut, fmt.Sprintln("str1:", str1+"\nstr2:", str2+"\nstr3:", str3+"\n------------------"))
	}

	if len(str1) > 0 {
		if str3[0] == str1[0] {
			if len(str1) == 1 {
				return true && shuffle("", str2, str3[1:])
			}
			return true && shuffle(str1[1:], str2, str3[1:])
		}
	}

	if len(str2) > 0 {
		if str3[0] == str2[0] {
			if len(str2) == 1 {
				return true && shuffle(str1, "", str3[1:])
			}
			return true && shuffle(str1, str2[1:], str3[1:])
		}
	}

	if len(str1) == 0 && len(str2) == 0 {
		return true
	}
	return false
}
