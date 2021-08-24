package singleton

import "errors"

func Execute(fn string, args ...int) int {
	switch fn {
	case "+":
		value, err := sum(args...)
		if err != nil {
			GetInstance().Error(err)
		}
		return value
	case "-":
		value, err := sub(args...)
		if err != nil {
			GetInstance().Error(err)
		}
		return value
	default:
		GetInstance().Error("Function not defined")
		return -1
	}
}

func sum(args ...int) (int, error) {
	if args == nil {
		return 0, errors.New("Invalid arguments")
	}
	result := 0
	for arg := range args {
		result += int(arg)
	}
	return result, nil
}

func sub(args ...int) (int, error) {
	if args == nil {
		return 0, errors.New("Invalid arguments")
	}
	result := int(args[0])
	for arg := range args[1:] {
		result -= int(arg)
	}
	return result, nil
}
