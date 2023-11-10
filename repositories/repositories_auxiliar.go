package repositories

import "log"

func Error(repository string, err error) {
	log.Fatalf("Error on '%s' repository: %v\n", repository, err)
}
