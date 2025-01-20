# Python BCC Example

## Install Dependencies

Please follow [this guide](https://github.com/iovisor/bcc/blob/master/INSTALL.md), to install all of the dependencies.

## Just Run

Using BCC, you simply define the kernel-side eBPF program as a string, while tasks like Clang compilation, loading, and attaching eBPF objects are done by the framework during runtime:

```
chmod +x main.py
sudo ./main.py
```
