package rules

import "reflect"
import (
	"fmt"
)

type Rule struct {}

func Invoke(name string, args... interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
		fmt.Println(args[i])
	}
	fmt.Printf( "Method Name: %s \n", name )
	reflect.ValueOf(Rule{}).MethodByName(name).Call(inputs)
}



