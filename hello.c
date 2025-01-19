#include <stdio.h>
#include <unistd.h>
#include <sys/resource.h>
#include <bpf/libbpf.h>
#include "hello.skel.h"

int main(void)
{
    struct hello_bpf *obj;
    int err = 0;

    struct rlimit rlim = {
        .rlim_cur = 512UL << 20,
        .rlim_max = 512UL << 20,
    };

    err = setrlimit(RLIMIT_MEMLOCK, &rlim);
    if (err) {
        fprintf(stderr, "failed to change rlimit\n");
        return 1;
    }

    obj = hello_bpf__open();
    if (!obj) {
        fprintf(stderr, "failed to open and/or load BPF object\n");
        return 1;
    }

    err = hello_bpf__load(obj);
    if (err) {
        fprintf(stderr, "failed to load BPF object %d\n", err);
        goto cleanup;
    }

    err = hello_bpf__attach(obj);
    if (err) {
        fprintf(stderr, "failed to attach BPF programs\n");
        goto cleanup;
    }

    // Infinite loop to keep the program running
    while (1) {
        sleep(1);
    }

cleanup:
    hello_bpf__destroy(obj);
    return err != 0;
}
