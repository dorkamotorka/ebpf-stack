package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpf hello hello.c

import (
	"os"
	"os/signal"
	"syscall"
	"log"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	// Load the compiled eBPF ELF and load it into the kernel.
	var objs helloObjects
	if err := loadHelloObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	// Attach Tracepoint
	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.HandleExecveTp, nil)
	if err != nil {
		log.Fatalf("Attaching Tracepoint: %s", err)
	}
	defer tp.Close()
	
	log.Println("eBPF program attached to tracepoint. Press Ctrl+C to exit.")

	// Set up signal handling to cleanly exit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Keep the program running until a signal is received
	<-stop

	log.Println("Received signal, exiting...")
}
