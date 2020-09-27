package peercred

import (
	"context"
	"net"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	l, err := net.Listen("unix", "@peercred-test")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		<-ctx.Done()
		conn.Close()
	}(ctx)

	conn, err := net.Dial("unix", "@peercred-test")
	if err != nil {
		t.Fatal(err)
	}

	cred, err := Read(conn.(*net.UnixConn))
	if err != nil {
		t.Fatalf("want nil err, got %v", err)
	}

	t.Logf("Cred: %+v", cred)

	pid, uid, gid := os.Getpid(), os.Getuid(), os.Getgid()
	if cred.PID != int32(pid) {
		t.Errorf("pid: want %d, got %d", pid, cred.PID)
	}
	if cred.UID != uint32(uid) {
		t.Errorf("uid: want %d, got %d", uid, cred.UID)
	}
	if cred.GID != uint32(gid) {
		t.Errorf("gid: want %d, got %d", gid, cred.GID)
	}
}
