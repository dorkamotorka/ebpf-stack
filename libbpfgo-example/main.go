package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"C"
	bpf "github.com/aquasecurity/tracee/libbpfgo"
)

func main() {
	// Create a channel to capture OS interrupt signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Load BPF object file
	bpfModule, err := bpf.NewModuleFromFile("hello.bpf.o")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading BPF module: %v\n", err)
		os.Exit(1)
	}
	defer bpfModule.Close() // Ensure the BPF module is closed on program exit

	// Load the BPF object into the kernel
	if err := bpfModule.BPFLoadObject(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading BPF object: %v\n", err)
		os.Exit(1)
	}

	// Get the BPF program from the loaded module
	prog, err := bpfModule.GetProgram("hello")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting BPF program: %v\n", err)
		os.Exit(1)
	}

	// Attach the BPF program to the kprobe for sys_execve
	if _, err := prog.AttachKprobe("__x64_sys_execve"); err != nil {
		fmt.Fprintf(os.Stderr, "Error attaching kprobe: %v\n", err)
		os.Exit(1)
	}

	// Wait for an interrupt signal to gracefully shutdown
	<-sig
	fmt.Println("Received interrupt signal, shutting down...")
}
