# Cilium ebpf Example

(bpf2go)[https://pkg.go.dev/github.com/cilium/ebpf/cmd/bpf2go] makes it extremely easy to compile and run eBPF program. You just have to run:
```
go build
go generate
sudo ./hello
```

by default, `bpf2go` internally uses `clang` with some helpful default flags. Among them are:

- `-g`: Includes debug information, which is necessary for BTF.

- `-strip llvm-strip`: Uses `llvm-strip` binary to reduce the size of the object file, as the `-g` flag adds DWARF debugging information that isn’t needed by eBPF programs.

- `-O2`: Ensures Clang produces BPF bytecode that passes the verifier. For example, Clang would normally output `callx <register>` for calling helper functions, but eBPF doesn’t support calling addresses from registers.

- `-target bpf`: Specifies the system the program is intended to run on (little-endian or big-endian). By default set to `bpf`, following the endianness of the CPU they’re compiled on.
