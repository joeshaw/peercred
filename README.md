# peercred

peercred is a Go package that wraps usage of the Linux `SO_PEERCRED`
socket option on Unix domain sockets.

From the [unix(7) man page](https://man7.org/linux/man-pages/man7/unix.7.html):

       SO_PEERCRED
              This read-only socket option returns the credentials of the
              peer process connected to this socket.  The returned creden‐
              tials are those that were in effect at the time of the call to
              connect(2) or socketpair(2).

              The argument to getsockopt(2) is a pointer to a ucred struc‐
              ture; define the _GNU_SOURCE feature test macro to obtain the
              definition of that structure from <sys/socket.h>.

              The use of this option is possible only for connected AF_UNIX
              stream sockets and for AF_UNIX stream and datagram socket
              pairs created using socketpair(2).

On Linux systems, the raw functionality is provided through the built-in `syscall` package and the `golang.org/x/sys/unix` package.  These packages, however, are not stable across operating systems and the usage of socket options is pretty low level.  This package encapsulates the functionality and returns errors on unsupported operating systems through an easy `Read` function.

The returned value provides the process ID, user ID, and group ID of the process on the other side of the Unix domain socket.  These values are populated by the Linux kernel and cannot be spoofed.  (However, these values are set at the time of socket creation and will not take into account privileges dropped afterward.)

## Usage

```go

conn, err := net.Dial("unix", "/var/run/somesocket")
if err != nil {
    log.Fatal(err)
}

cred, err := peercred.Read(conn.(*net.UnixConn))
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", cred)
// => &{PID:2002 UID:1000 GID:1000}
```

## License

Copyright 2020 Joe Shaw

`peercred` is licensed under the MIT license.  See the LICENSE file for details.