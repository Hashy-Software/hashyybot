package main

import (
	"fmt"
	"reflect"

	command "github.com/Hashy-Software/hashyybot/pkg/command"
)

func main() {
	myCommandHandler := func(arg1 int, arg2 string, arg3 int) error {
		fmt.Printf("Executing my command with arguments: %d, %s, %d\n", arg1, arg2, arg3)
		return nil
	}

	argTypes := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(""), reflect.TypeOf(0)}
	keywordArgs := make(map[string]interface{})
	myCommand := command.NewCommand("mycommand", myCommandHandler, argTypes, keywordArgs)

	err := myCommand.Execute(42, "hello", 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
