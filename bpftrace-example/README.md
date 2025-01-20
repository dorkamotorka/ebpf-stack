# Bpftrace example

## Install Dependencies

Please follow [this guide](https://github.com/bpftrace/bpftrace/blob/master/INSTALL.md) to install the dependencies.

## Just Run

- One Liner:
```
sudo bpftrace -e 'tracepoint:syscalls:sys_enter_execve { printf("Hello world\n"); }'
```

- Or run the script using `sudo bpftrace hello.bt`
