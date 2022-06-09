package znodes

import (
	"errors"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
)

func CreatePathIfNecessary(conn *zk.Conn, path string) error {
	return CreatePathWithDataIfNecessary(conn, path, false, nil)
}

func CreatePathWithDataIfNecessary(conn *zk.Conn, path string, ephemeral bool, data []byte) error {

	if len(path) <= 1 {
		return errors.New("zk path is empty")
	}
	var segs []string
	if strings.HasPrefix(path, "/") {
		segs = strings.Split(path[1:], "/")
	} else {
		segs = strings.Split(path, "/")
	}
	p := ""
	for idx, s := range segs {
		p += "/" + s
		existed, _, _ := conn.Exists(p)
		var err error
		if !existed && idx == len(segs)-1 {
			if ephemeral {
				_, err = conn.Create(p, data, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
			} else {
				_, err = conn.Create(p, data, 0, zk.WorldACL(zk.PermAll))
			}
		} else if !existed {
			_, err = conn.Create(p, nil, 0, zk.WorldACL(zk.PermAll))
		}
		if err != nil {
			return err
		}
	}
	return nil
}
