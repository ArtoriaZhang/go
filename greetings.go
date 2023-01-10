package greetings

import "fmt"
import "errors"

func Hello(name string) (string, error) {
	if (name == "") {
		return "", errors.New("Empty name");
	}
	message := fmt.Sprintf("Hi, %v. Welcom!", name)
	return message, nil;
}
