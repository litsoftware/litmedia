package timeh

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func In(t time.Time, zone string) (time.Time, error) {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return t, err
	}

	return t.UTC().In(loc), nil
}

func InPbTimestamp(t time.Time, zone string) (*timestamppb.Timestamp, error) {
	if t.IsZero() {
		return &timestamppb.Timestamp{}, nil
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		return nil, err
	}

	_, offset := t.In(loc).Zone()
	return &timestamppb.Timestamp{Seconds: t.Unix() + int64(offset)}, nil
}
