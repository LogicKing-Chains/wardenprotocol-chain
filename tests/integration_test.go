package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/warden-protocol/wardenprotocol/tests/cases"
	"github.com/warden-protocol/wardenprotocol/tests/framework"
)

var list bool

func TestMain(m *testing.M) {
	flag.BoolVar(&list, "list-only", false, "If set, prints the list of the available test cases instead of executing them.")
	flag.Parse()
	os.Exit(m.Run())
}

func TestIntegration(t *testing.T) {
	var casesToRun = cases.List()
	if list {
		fmt.Println("Available test cases:")
		for _, c := range casesToRun {
			fmt.Println("*", getName(c))
		}
		return
	}

	build := framework.Build(t)
	startTime := time.Now()

	for _, c := range casesToRun {
		t.Run(getName(c), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			fmt.Println("setup test:", getName(c))
			c.Setup(t, ctx, build)
			fmt.Println("run test:", getName(c))
			c.Run(t, ctx, build)
		})
	}

	fmt.Println("tests duration:", time.Since(startTime))
}

func getName(v any) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
