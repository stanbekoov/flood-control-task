package floodcontrol_test

import (
	"context"
	"log"
	"task/config"
	"task/floodcheck"
	"testing"
)

func TestCheck(t *testing.T) {
	want := []bool{true, true, true, true, true, false, false}
	config := config.New()
	config.Flood.K = 5
	config.Flood.N = 5
	fc := floodcheck.New(config)
	ctx := context.Background()
	for i := 0; i < len(want); i++ {
		get, err := fc.Check(ctx, 0)
		if err != nil {
			log.Fatal(err)
		}
		if get != want[i] {
			t.Fatalf("for i = %d, res == %t, want == %t", i, get, want[i])
		}
	}
}
