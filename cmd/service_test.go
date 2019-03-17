package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"testing"
	_ "unsafe"
)

//go:linkname getMapTags github.com/grpc-ecosystem/go-grpc-middleware/tags.newTags
func getMapTags() grpc_ctxtags.Tags

//go:linkname setInContext github.com/grpc-ecosystem/go-grpc-middleware/tags.setInContext
func setInContext(ctx context.Context, tags grpc_ctxtags.Tags) context.Context

func TestUnaryServerMetdataTagInterceptor(t *testing.T) {

	tags := getMapTags()
	tags.Set("test", "YOYO")

	ctx := context.Background()

	newCtx := setInContext(ctx, tags)

	fmt.Printf("ctx: %+v, tags: %+v\n", newCtx, tags)

	tags = grpc_ctxtags.Extract(newCtx)
	tags.Set("userID", "yoyoyousei")

	fmt.Printf("ctx: %+v, tags: %+v\n", newCtx, tags)
}
