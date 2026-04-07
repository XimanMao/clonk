// main.go
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var listRoasts bool

func init() {
	flag.BoolVar(&listRoasts, "list", false, "print all roast messages and exit")
}

func main() {
	flag.Parse()

	if listRoasts {
		for _, r := range roasts {
			fmt.Println(r)
		}
		os.Exit(0)
	}

	if os.Geteuid() != 0 {
		fmt.Fprintln(os.Stderr, "clonk requires root privileges for accelerometer access.\nRun with: sudo clonk")
		os.Exit(1)
	}

	fmt.Println("🔨 clonk: listening for bonks... (ctrl+c to quit)")
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "clonk: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	bonkCh := make(chan struct{}, 1)

	// Start accelerometer monitor
	errCh := make(chan error, 1)
	go func() {
		errCh <- monitorAccelerometer(ctx, bonkCh)
	}()

	// Listen for bonk events and dispatch roasts
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n👋 bye!")
			return nil
		case err := <-errCh:
			return err
		case <-bonkCh:
			msg := randomRoast()
			if err := dispatchRoast(msg); err != nil {
				fmt.Fprintf(os.Stderr, "warning: %v\n", err)
			}
		}
	}
}
