package main

import (
	"bookstore/pb"
	"context"
	"image/color/palette"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

const (
	defaultCursor = 0
	defaultPageSize = 1 
)
type server struct {
	pb.UnimplementedBookstoreServer

	bs *bookstore // gRPC服务
}

// 列出所有书架的RPC方法
func (s *server)  ListShelves(ctx context.Context, in *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	//调用gorm 的方法
	sl,err:=s.bs.ListShelves(ctx)
	if err == gorm.ErrEmptySlice{
		return &pb.ListShelvesResponse{},nil
	}
	if err!=nil{ // 查询数据库出错
		return nil,status.Error(codes.Internal,"qurey failed")
	}
	nsl := make([] *pb.Shelf,0,len(sl))
	for _,s := range sl {
		nsl = append(nsl, &pb.Shelf{
			Id:    s.ID,
			Theme: s.Theme,
			Size:  s.Size,
		})
	}
	return &pb.ListShelvesResponse{Shelves: nsl},nil
}

// 创建书架
func (s *server) CreateShelf(ctx context.Context,in *pb.CreateShelfRequest) (*pb.Shelf,error){
	//参数检查
	if len(in.GetShelf().GetTheme()) == 0{
		return nil,status.Error(codes.InvalidArgument,"invalid theme")
	}
	// 准备数据
	data := Shelf{
		Theme: in.GetShelf().GetTheme(),
		Size: in.GetShelf().GetSize(),
	}

	// 去数据库创建
	ns ,err := s.bs.CreateShelf(ctx,data)
	if err!=nil{
		return nil,status.Error(codes.Internal,"create failed")
	}
	return &pb.Shelf{Id: ns.ID,Theme: ns.Theme,Size: ns.Size},nil
}

// GetShelf 根据id获取书架
func (s *server) GetShelf (ctx context.Context,in *pb.GetShelfRequest) (*pb.Shelf,error){
	//参数检查
	if in.GetShelf()<=0{
		return nil,status.Error(codes.InvalidArgument,"getshelf failed")
	}
	//查询数据库
	shelf,err:= s.bs.GetShelf(ctx,in.GetShelf())
	if err!=nil{
		return nil,status.Error(codes.Internal,"query failed")
	}
	
	// 封装返回数据
	return &pb.Shelf{
		Id:    shelf.ID,
		Theme: shelf.Theme,
		Size:  shelf.Size,
	},nil
}

// DeleteShelf 删除书架
func (s *server) DeleteShelf (ctx context.Context,in *pb.DeleteShelfRequest)(*empty.Empty,error){
	//参数检测
	if in.GetShelf()<=0{
		return nil,status.Error(codes.InvalidArgument,"invlid shelf id")
	}
	// 查询数据库
	err := s.bs.DeleteShelf(ctx,in.GetShelf())
	if err!=nil{
		return nil,status.Error(codes.Internal,"delete failed")
	}
	return &empty.Empty{},nil
}

// 列出所有图书
func (s *server) ListBooks(ctx context.Context,in *pb.ListBooksRequest)(*pb.ListBooksResponse,error){
	// check 参数 (shelf_id,token:是否存在，token:是否有效：三个字段值)
	var (

	)
	if in.GetShelf()<=0{
		return nil,status.Error(codes.InvalidArgument,"invalid shelf_id")
	}
	if in.GetPageToken() == ""{
		// 获取到的token为空值,没有分页token默认第一页
		

	}else{
		// 判断是否有分页
		if len(in.GetPageToken())>0{
			// 先转成token再dencode
			pageInfo := Token(in.GetPageToken()).Decode()
			//判断结果是否有效(三个字段)
			if pageInfo.NextID == "" || pageInfo.NextTimeAtUTC == 0 || pageInfo.NextTimeAtUTC > time.Now().Unix() || pageInfo.PageSize <= 0 {e
				return nil,status.Error(codes.InvalidArgument,"invalid pagetoken")
			}
			cursor = pageInfo.NextID
			pageSize = pageInfo.PageSize
	}

	// 查询数据库，基于游标(getbyID,pagesize+1：判断是否存在下一页)
	// 差的条目会多，需要定义真实请求需要的数目,
	bookList,err:= s.bs.GetBookListByID(ctx,in.GetShelf(),cursor,pageSize+1){
		if err != nil {
			fmt.Printf("GetBookListByShelfID failed, err:%v\n", err)
			return nil, status.Error(codes.Internal, "query failed")
		}
		// 如果查询出来的结果比pageSize大，那么就说明有下一页
		var (
			hasNextPage bool
			nextPageToken string
			realSize int=(bookList)
		)
		if lne(bookList) > pageSize {
			hasNextPage = true
			realSize = pageSize
		}
	}

	// 封装返回
	// []*Books -> [] *pb.Books
	res := make([]*pb.Book,0,len(pageList))
	for i=0;i<realSize;i++{
		res = append(res, &pb.Book{
			Id:     bookList[i].ID,
			Author: bookList[i].Author,
			Tite:   bookList[id].Tite,
		})
	}

	// 如果有下一页返回下一页的page_token
	if hasNextPage {
		nextPageInfo:Page{
			NextID: strconv.FormatInt(res[len(realSize-1)].Id),
			PageSize: int64(pageSize),
		}
	}
	nextPageToken = string (nextPageInfo.Encode())

	return &pb.ListBooksResponse{Books:res,NextPageToken: nextPageToken},nil
}