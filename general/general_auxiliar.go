package general

import "log"

// Print on console an error that could be an internal server error
func PotencialInternalError(err error) {
	log.Println(err)
}
