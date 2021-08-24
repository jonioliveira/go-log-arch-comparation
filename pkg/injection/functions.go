package injection

import (
	"errors"

	"github.com/sirupsen/logrus"
)

func Execute(log *logrus.Logger, fn string, args ...int) int {
	switch fn {
	case "+":
		value, err := sum(args...)
		if err != nil {
			log.Error(err)
		}
		return value
	case "-":
		value, err := sub(args...)
		if err != nil {
			log.Error(err)
		}
		return value
	default:
		log.Error("Function not defined")
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
