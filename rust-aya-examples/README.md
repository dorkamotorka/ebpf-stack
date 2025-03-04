# Rust Aya Example

There are many great resources, for developing eBPF programs using Aya that do much better job explaining the concepts than I ever could. Just to reference a few: 

- [eBPF Aya Book](https://aya-rs.dev/book/)
- [Examples of small eBPF projects based on Aya](https://github.com/vadorovsky/aya-examples/tree/main)
- [Head First eBPF Aya Example](https://www.ebpf.top/en/post/ebpf_rust_aya/)

To quickly run an example, do:
```
git clone https://github.com/littlejo/aya-examples
cd aya-examples/tracepoint-binary
RUST_LOG=info cargo run
```
**NOTE**: It might take a while to install all the dependencies.
