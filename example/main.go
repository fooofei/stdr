package main

import (
	stdlog "log"
	"os"

	"github.com/go-logr/stdr"
)

type E struct {
	str string
}

func (e E) Error() string {
	return e.str
}

func main() {
	stdr.SetVerbosity(1)
	log := stdr.New(stdlog.New(os.Stderr, "", stdlog.LstdFlags|stdlog.Lshortfile))
	log = log.WithName("MyName").WithValues("user", "you")
	log.Info("hello", "val1", 1, "val2", map[string]int{"k": 1})
	log.V(1).Info("you should see this")
	log.V(3).Info("you should NOT see this")
	log.Error(nil, "uh oh", "trouble", true, "reasons", []float64{0.1, 0.11, 3.14})
	log.Error(E{"an error occurred"}, "goodbye", "code", -1)
}
