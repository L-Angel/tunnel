package json

import (
	"fmt"
)

func Main() {
}

func Marshal(v interface{}) ([]byte, error) {
	fmt.Println(v)
	return nil, nil
}

func UnMarshal(v []byte, inst interface{}) error {
	fmt.Println(v, inst)
	return nil
}
