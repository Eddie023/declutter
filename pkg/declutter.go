package declutter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Tidy(path string) {
	if err := run(); err != nil {
		log.Error("Failed with err", err)
	}
}

func run() error {

	// =========================
	// Configuration

	cfg := struct {
		Foo string
	}{
		Foo: "apple",
	}

	fmt.Println("the cfg", cfg)

	return nil
}
