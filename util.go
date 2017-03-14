package apiai

import (
	"fmt"
	"time"

	"github.com/jmcvetta/randutil"
)

func randomString() string {
	rand, _ := randutil.String(12, randutil.Alphabet)
	ts := time.Now().Unix()
	return fmt.Sprintf("%v%v", rand, ts)
}
