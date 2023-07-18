package command

import (
	"errors"
	"reflect"
)

type Command struct {
	name        string
	callback    interface{}
	argTypes    []reflect.Type
	keywordArgs map[string]interface{}
}

func NewCommand(name string, callback interface{}, argTypes []reflect.Type, keywordArgs map[string]interface{}) *Command {
	return &Command{
		name:        name,
		callback:    callback,
		argTypes:    argTypes,
		keywordArgs: keywordArgs,
	}
}

func (c *Command) Execute(args ...interface{}) error {
	if len(args) != len(c.argTypes) {
		return errors.New("Invalid argument count")
	}

	for i, arg := range args {
		if argType := c.argTypes[i]; !argType.AssignableTo(reflect.TypeOf(arg)) {
			return errors.New("Invalid argument type")
		}
	}

	returnValue := reflect.ValueOf(c.callback).Call(convertArgs(args))
	if len(returnValue) > 0 && !returnValue[0].IsNil() {
		return returnValue[0].Interface().(error)
	}

	return nil
}

func convertArgs(args []interface{}) []reflect.Value {
	convertedArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		convertedArgs[i] = reflect.ValueOf(arg)
	}
	return convertedArgs
}
