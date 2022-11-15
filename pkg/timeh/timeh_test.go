package timeh

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestInPbTimestamp(t *testing.T) {
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Shanghai")

	fmt.Println(now)
	fmt.Println(now.UTC())
	fmt.Println(now.UTC().In(loc))
	fmt.Println("----------------")
	fmt.Println(now.Unix())
	fmt.Println(now.UTC().Unix())
	fmt.Println(now.UTC().In(loc).Unix())
	fmt.Println("----------------")
	name, offset := now.In(loc).Zone()
	fmt.Println(name, offset)
	fmt.Println(timestamppb.New(now.In(loc)))
	fmt.Println(&timestamppb.Timestamp{Seconds: now.Unix() + int64(offset)})
}
