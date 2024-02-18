package pets

import "math/rand"

func generateBool() bool {
	return rand.Intn(2) == 0
}
