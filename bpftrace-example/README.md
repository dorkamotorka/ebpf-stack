# Bpftrace example

- One Liner:
```
sudo bpftrace -e 'tracepoint:syscalls:sys_enter_execve { printf("Hello world\n"); }'
```

- Or run the script using `sudo bpftrace hello.bt`
