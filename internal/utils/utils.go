package utils

import "fmt"

func GenerateFlag(stage int) string {
	return fmt.Sprintf("FLAG-%d-MOD-CTF", stage)
}
