package modules

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nombre en blanco")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %s! Bienvenido",
		"¡%s! ¿Cómo estás?",
		"¡%s! ¿Qué tal tu día?",
		"¡Que bueno verte, %s!",
	}

	return formats[rand.Intn(len(formats))]
}
