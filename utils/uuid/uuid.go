package uuid

import "github.com/google/uuid"

func Generate() (string, error) {
	b, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	id := b.String()
	return id, nil
}
