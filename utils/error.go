package utils

import "fmt"

func ModuleError() {
	fmt.Printf("correct module is required")
}

func LoginError() {
	fmt.Printf("please set MIKRUS_SRV and MIKRUS_KEY correctly")
}

func NoCommandProvidedError() {
	fmt.Printf("please provide command to execute")
}
