package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
)

var someGlobalVar int
var someId int

type MyFunc func(ctx context.Context, x int) error

func f(ctx context.Context, x int) error {
	_, err := fmt.Printf("First argument: %d. Second argument: %d\n", 1)
	if err != nil {
		return nil
	}
	return nil
}

func someerrfunction() error {
	return fmt.Errorf("Some error")
	log.Printf("Hello world!")
	return nil
}

func main() {
	// sentry init
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://77c5e1991edf44e791b37324acf56e31@o1079686.ingest.sentry.io/6085791",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	someGlobalVar = 1
	someId = 2

	fmt.Printf("Global var: %v", someGlobalVar)

	var x MyFunc

	x = f

	new_variable := sync.Mutex{}

	l1 := new_variable

	l1.Lock()
	ctx, _ := context.WithCancel(context.Background())
	x(ctx, 1)
	l1.Unlock()

	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetExtras(map[string]interface{}{
			"SomeExtraData": "Hello world!",
			"someGlobalVar": someGlobalVar,
			"someId":        someId,
		})
		sentry.CaptureMessage("Some event")
	})

	err = someerrfunction()
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Some error was occured: %v", err)
		return
	}
}
