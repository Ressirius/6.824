package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"

//
// example to show how to declare the arguments
// and reply for an RPC.
//

// 创建任务枚举
type MRJobType int

const (
    UnknownJob MRJobType = iota     // 未知任务
    MapJob                          // map任务
    ReduceJob                       // reduce任务
    WaitJob                         // map任务分配完闭，worker等待reduce任务
)

// 任务状态枚举
type MRJobState int

const (
    UnknownState MRJobState = iota  // 未知状态
    UnFinish                        // 未完成
    Progressing                     // 进行中
    Finish                          // 完成
)

type Job struct {
    JobType MRJobType       // 任务类型
    JobState MRJobState     // 任务状态

    JobId int               // 任务键
    OperateFileName string  // 任务操纵的文件
}

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.


// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
