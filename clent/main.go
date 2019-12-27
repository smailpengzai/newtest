//package main
//
//import (
//	"flag"
//	pb "grpc_test/protoDir"
//	"time"
//	"github.com/astaxie/beego"
//	"golang.org/x/net/context"
//	"io"
//	"math/rand"
//	"google.golang.org/grpc/credentials"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/testdata"
//	"fmt"
//)
//
//var (
//	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
//	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
//	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
//	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
//)
//
//func printFeature(client pb.RouteGuideClient, point *pb.Point) {
//	beego.Notice(fmt.Sprintf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude))
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	feature, err := client.GetFeature(ctx, point)
//	if err != nil {
//		beego.Error(fmt.Sprintf("%v.GetFeatures(_) = _, %v: ", client, err))
//	}
//	beego.Notice(feature)
//}
//
//func printFeatures(client pb.RouteGuideClient, rect *pb.Rectangle) {
//	beego.Notice(fmt.Sprintf("Looking for features within %v", rect))
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	stream, err := client.ListFeatures(ctx, rect)
//	if err != nil {
//		beego.Error(fmt.Sprintf("%v.ListFeatures(_) = _, %v", client, err))
//	}
//	for {
//		feature, err := stream.Recv()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			beego.Error(fmt.Sprintf("%v.ListFeatures(_) = _, %v", client, err))
//		}
//		beego.Notice(feature)
//	}
//}
//
//func runRecordRoute(client pb.RouteGuideClient) {
//	// Create a random number of random points
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
//	var points []*pb.Point
//	for i := 0; i < pointCount; i++ {
//		points = append(points, randomPoint(r))
//	}
//	beego.Notice(fmt.Sprintf("Traversing %d points.", len(points)))
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	stream, err := client.RecordRoute(ctx)
//	if err != nil {
//		beego.Error(fmt.Sprintf("%v.RecordRoute(_) = _, %v", client, err))
//	}
//	for _, point := range points {
//		if err := stream.Send(point); err != nil {
//			beego.Error(fmt.Sprintf("%v.Send(%v) = %v", stream, point, err))
//		}
//	}
//	reply, err := stream.CloseAndRecv()
//	if err != nil {
//		beego.Error(fmt.Sprintf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil))
//	}
//	beego.Notice(fmt.Sprintf("Route summary: %v", reply))
//}
//func randomPoint(r *rand.Rand) *pb.Point {
//	lat := (r.Int31n(180) - 90) * 1e7
//	long := (r.Int31n(360) - 180) * 1e7
//	return &pb.Point{Latitude: lat, Longitude: long}
//}
//
//func runRouteChat(client pb.RouteGuideClient) {
//	notes := []*pb.RouteNote{
//		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
//		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
//		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
//		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
//		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
//		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	stream, err := client.RouteChat(ctx)
//	if err != nil {
//		beego.Error(fmt.Sprintf("%v.RouteChat(_) = _, %v", client, err))
//	}
//	waitc := make(chan struct{})
//	go func() {
//		for {
//			in, err := stream.Recv()
//			if err == io.EOF {
//				// read done.
//				close(waitc)
//				return
//			}
//			if err != nil {
//				beego.Error(fmt.Sprintf("Failed to receive a note : %v", err))
//			}
//			beego.Notice(fmt.Sprintf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude))
//		}
//	}()
//	for _, note := range notes {
//		if err := stream.Send(note); err != nil {
//			beego.Error(fmt.Sprintf("Failed to send a note: %v", err))
//		}
//	}
//	stream.CloseSend()
//	<-waitc
//}
//
//func main() {
//	flag.Parse()
//	var opts []grpc.DialOption
//	if *tls {
//		if *caFile == "" {
//			*caFile = testdata.Path("ca.pem")
//		}
//		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
//		if err != nil {
//			beego.Error(fmt.Sprintf("Failed to create TLS credentials %v", err))
//		}
//		opts = append(opts, grpc.WithTransportCredentials(creds))
//	} else {
//		opts = append(opts, grpc.WithInsecure())
//	}
//	conn, err := grpc.Dial(*serverAddr, opts...)
//	if err != nil {
//		beego.Error(fmt.Sprintf("fail to dial: %v", err))
//	}
//	defer conn.Close()
//	client := pb.NewRouteGuideClient(conn)
//
//	// Looking for a valid feature
//	printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
//
//	// Feature missing.
//	printFeature(client, &pb.Point{Latitude: 0, Longitude: 0})
//
//	// Looking for features between 40, -75 and 42, -73.
//	printFeatures(client, &pb.Rectangle{
//		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
//		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
//	})
//
//	// RecordRoute
//	runRecordRoute(client)
//
//	// RouteChat
//	runRouteChat(client)
//}

package main

import "fmt"

func main() {
	x := "root@lps-web-5b97889d56-k4jnk:/usr/local/tomcat/logs# tail -fn 200 catalina.2019-10-05.log"
	fmt.Print(len([]byte(x)))
}