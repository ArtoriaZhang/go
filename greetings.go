package greetings

import ("fmt"
	"errors"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if (name == "") {
		return "", errors.New("Empty name");
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil;
}

func Hellos(name []string) (map[string]string, error) {
	resultMap := make(map[string]string);

	for _, e := range name {
		msg, err := Hello(e);

		if (err != nil) {
			return nil, err;
		}

		resultMap[e] = msg;
	}

	return resultMap, nil;
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string{
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v!. Well met",
		"High five, %v",
	};

	return formats[rand.Intn(len(formats))];
}
