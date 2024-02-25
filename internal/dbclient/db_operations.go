package dbclient

import (
	"context"
	"fmt"

	merr "github.com/autotest-plan/errors"
	mlog "github.com/autotest-plan/log"
	"github.com/autotest-plan/rpcdefine/go/dbadapter"
	"github.com/autotest-plan/rpcdefine/go/message"
	"google.golang.org/grpc"
)

type DbOperations struct {
	conn   *grpc.ClientConn
	client dbadapter.DBAdapterClient
	*mlog.Logger
}

func NewDbOperations(ctx context.Context, ip string, port int, logger *mlog.Logger) (*DbOperations, error) {
	var opts []grpc.DialOption

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), opts...)
	if err != nil {
		logger.Errorf("db客户端grpc拨号失败: %s\n", err.Error())
		return nil, err
	}
	client := dbadapter.NewDBAdapterClient(conn)

	return &DbOperations{conn: conn, client: client, Logger: logger}, nil
}

func (dbo *DbOperations) Store(tasks []*message.Task) error {
	result, err := dbo.client.Store(context.TODO(), &message.Tasks{Tasks: tasks})
	if err != nil || !result.Result {
		message := fmt.Sprintf("存储任务失败\nresult: %v\nerr: %+v\n", result.Result, err)
		dbo.Errorf(message)
		return merr.Error(dbadapter.DBCode, message)
	}
	return nil
}

func (dbo *DbOperations) LoadFailedTasks() *message.Tasks {
	// TODO: Filter中value的值限定了string，应该使用更广泛的类型
	tasks, err := dbo.client.LoadSorted(context.TODO(), &message.Filter{Kv: map[string]string{"result": "false"}})
	if err != nil {
		dbo.Errorf("加载失败的任务失败:\n%s\n", err.Error())
		return nil
	}
	return tasks
}

func (dbo *DbOperations) LoadSuccessTasks() *message.Tasks {
	// TODO: Filter中value的值限定了string，应该使用更广泛的类型
	tasks, err := dbo.client.LoadSorted(context.TODO(), &message.Filter{Kv: map[string]string{"result": "true"}})
	if err != nil {
		dbo.Errorf("加载失败的任务失败:\n%s\n", err.Error())
		return nil
	}
	return tasks
}
