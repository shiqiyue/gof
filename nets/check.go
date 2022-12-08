package nets

import (
	"github.com/shiqiyue/gof/ferror"
	"net"
	"time"
)

func TcpCheck(address string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", address, timeout)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	if err != nil {
		return err
	}
	if conn != nil {
		return nil
	} else {
		return ferror.New("connect fail")
	}
}
