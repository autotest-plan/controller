package listener

import (
	"context"
	"net/http"
	"sync"

	"github.com/autotest-plan/rpcdefine/go/dbadapter"
	"github.com/autotest-plan/rpcdefine/go/executor"
	"github.com/autotest-plan/rpcdefine/go/message"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type triggerHandler struct {
	sync.RWMutex
	ctx  context.Context
	exec executor.ExecutorClient
	db   dbadapter.DBAdapterClient
}

func newTriggerHandler(ctx context.Context) (*triggerHandler, error) {
	execConn, err := grpc.Dial("executorIP:Port", grpc.WithDefaultCallOptions())
	if err != nil {
		return nil, err
	}
	dbConn, err := grpc.Dial("dbAdapterIP:Port", grpc.WithDefaultCallOptions())
	if err != nil {
		return nil, err
	}
	return &triggerHandler{
		ctx:  ctx,
		exec: executor.NewExecutorClient(execConn),
		db:   dbadapter.NewDBAdapterClient(dbConn),
	}, nil
}

func (t *triggerHandler) Trigger(c *gin.Context) {
	t.Lock()
	defer t.Unlock()

	// 获取trigger中包含的任务
	var tasks message.Tasks
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: tasks对比数据库中的任务决定是否读取
	if tasks.Tasks == nil || len(tasks.Tasks) == 0 {
		// TODO: 加载哪些任务
		t.db.LoadSorted(t.ctx, &message.Filter{}, grpc.EmptyCallOption{})
	}

	// TODO: tasks最终要递交到executor
	t.exec.Execute(t.ctx, tasks.Tasks[0], grpc.EmptyCallOption{})
}
