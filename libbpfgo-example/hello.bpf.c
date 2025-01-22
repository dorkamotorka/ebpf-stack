//go:build ignore
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

char LICENSE[] SEC("license") = "Dual BSD/GPL";

SEC("kprobe/sys_execve")
int hello(void *ctx) {
    bpf_printk("Hello world");
    return 0;
}
