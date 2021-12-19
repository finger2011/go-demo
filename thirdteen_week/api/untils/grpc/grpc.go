package grpc

import (
	pb "api/grpc/service"
	err "api/untils/error"
	"api/untils/log"
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

type GRPCTestServer struct {
	pb.UnimplementedChatServer
}

func (s *GRPCTestServer) Test(ctx context.Context, req *pb.Request) (*pb.Response, error) { //实现服务端收到客户端请求后的应答方法
	log.Logger.Debug(fmt.Sprintln("客户端传来：", req.GetName()))
	return &pb.Response{Message: "服务端应答---->客户端传来的内容：" + req.GetName()}, nil //返回给客户端的应答内容 req.GetName()就是客户端发来的名字，也就是我们在proto文件中定义Request中的 string name = 1;
}

//TODO:不同的GRPCTestServer应该当参数传入
func RunGRPC(addr string) *err.ApiError {
	l, e := net.Listen("tcp", addr)
	if e != nil {
		return err.NewGRPCError(e.Error())
	}
	s := grpc.NewServer()                       // NewServer创建一个gRPC服务
	pb.RegisterChatServer(s, &GRPCTestServer{}) //服务注册
	if e = s.Serve(l); e != nil {               //启动GRPC服务
		return err.NewGRPCError(e.Error())
	}
	log.Logger.Debug(fmt.Sprintln("启动GRPC服务成功：", addr))
	return nil
}

//TODO: call addr, method string 参数修改
func CallGRPC(addr string) (*pb.Response, error) {
	conn, e := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock()) // 用grpc去连接服务端，grpc.WithInsecure() - 不需要传入证书，grpc.WithBlock() - 让客户端进入连接状态返回连接套接字
	//需要证书：
	//tlsInsecure, _ := credentials.NewClientTLSFromFile("xxx.crt", “服务器名”)
	//conn , err := grpc.Dial(":8001",grpc.WithTransportCredentials(tlsInsecure))
	if e != nil {
		log.Logger.Debug(fmt.Sprintln("无法连接客户端:", addr, "; err", e))
		return nil, err.NewGRPCError(e.Error())
	}
	defer conn.Close()
	c := pb.NewChatClient(conn)                                           //调用方法创建一个客户端连接
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) //设置上下文超时取消
	defer cancel()
	resp, e := c.Test(ctx, &pb.Request{ //开始向服务端发送数据，数据来源是cmd命令行输入的内容
		Name: os.Args[1],
	})
	if e != nil {
		log.Logger.Debug(fmt.Sprintln("没有得到响应", e))
		return nil, err.NewGRPCError(e.Error())
	}
	log.Logger.Debug(fmt.Sprintln("获取返回：", resp.GetMessage()))
	return resp, nil
}
