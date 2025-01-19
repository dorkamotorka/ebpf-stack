# Libbpf Hello World Example

## Prerequisites

In order to run this example on **Ubuntu/Debian**, you need to install the following dependencies:
```
sudo apt update && sudo apt install -y clang llvm libelf-dev zlib1g-dev make linux-tools-common linux-tools-generic linux-tools-$(uname -r)
```

We also need libbpf, which is added a submodule to this repository. You need to go inside the `libbpf/src` and trigger:
```
make BUILD_STATIC_ONLY=1 OBJDIR=../build/libbpf DESTDIR=../build INCLUDEDIR= LIBDIR= UAPIDIR= install
```

## Compile and Run

For educational purposes the project doesn't use a `Makefile`, which you'll often see in other projects. 

Instead we manually build it using the following commands:
- Compile the eBPF program:
```
clang -g -O2 -target bpf -c hello.bpf.c -o hello.bpf.o
```

- Generate an eBPF Skeleton (The skeleton file simplifies the process for the user space programs to access global variables and work with BPF programs):
```
bpftool gen skeleton hello.bpf.o > hello.skel.h
```

- Compile the user-space program and link it with libbpf library:
```
clang -g -O2 -Wall hello.c libbpf/build/libbpf.a -lelf -lz -o hello
```
