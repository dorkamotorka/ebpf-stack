# Libbpfgo Example

## Install dependencies

```
sudo apt-get update
sudo apt-get install libbpf-dev make clang llvm libelf-dev
```

## Compile and Run

First compile it using `clang`:
```
clang -O2 -c -target bpf -o hello.bpf.o hello.bpf.c
CC=gcc CGO_CFLAGS="-I /usr/include/bpf" CGO_LDFLAGS="/usr/lib/x86_64-linux-gnu/libbpf.a" go build -o hello
```

Then you can run it using:
```
sudo ./hello
```
