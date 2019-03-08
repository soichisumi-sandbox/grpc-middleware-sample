package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/soichisumi-sandbox/grpc-middleware-sample/api-pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"reflect"
	"strings"
)

const (
	port             = ":3000"
	AuthorizationKey = "authorization"
)

// UnaryServerMetdataTagInterceptor ...
func UnaryServerMetdataTagInterceptor(fields ...string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if ctxMd, ok := metadata.FromIncomingContext(ctx); ok {
			tags := grpc_ctxtags.Extract(ctx)
			fmt.Printf("ctxMd: %+v\n", ctxMd)
			fmt.Printf("tags: %+v\n", tags)

			//for _, field := range fields {
			//	if values, present := ctxMd[field]; present {
			//		tags.Set(field, strings.Join(values, ","))
			//	}
			//}
		}
		return handler(ctx, req)
	}
}

// NOTE: using reflect package may affect to performance
// NOTE: converted map is shallow. nested struct cannot be retrieved from resMap.
func structToMap(target interface{}) map[string]interface{} {
	resMap := make(map[string]interface{})
	elem := reflect.ValueOf(target).Elem()
	for i := 0; i < elem.NumField(); i++ {
		//field := elem.Type().Field(i).Tag.Get("json")
		field := strings.ToLower(elem.Type().Field(i).Name)
		value := elem.Field(i).Interface()
		resMap[field] = value
	}
	return resMap
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %+v\n", err)
	}

	server, err := NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %+v\n", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				UnaryServerMetdataTagInterceptor(),
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(func(fullmethod string, req interface{}) map[string]interface{} { // ex. type of req: apipb.AddUser
					fmt.Printf("extractor. fullmethod: %s\n, req: %+v\n", fullmethod, req)

					reqMap := structToMap(req)
					fmt.Printf("map: %+v\n", reqMap)
					return nil
				})),
			)))

	apipb.RegisterUserServiceServer(s, server)
	reflection.Register(s)
	fmt.Printf("grpc server is running on port:%s...\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
