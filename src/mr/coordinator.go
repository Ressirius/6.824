package mr

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"  // 锁
)


type Coordinator struct {
	// Your definitions here.
    files []string              // 操作的文件
    nReduce int                 // reduce任务数
    allJob map[MRJobType][]Job  // 所有的任务清单
    
    // map job
    mapJobFinish bool           // map任务是否全部完成
    mapJob []Job                // map任务清单

    // reduce job
    reduceJobFinish bool        // reduce任务是否全部完成
    reduceJob []Job             // reduce任务清单

    // lock
    lock sync.Mutex             // 锁
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}


//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.


	return ret
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.
    

    // 初始化Coordinator
    c.nReduce = nReduce
    c.files = files

    c.mapJobFinish = false
    c.reduceJobFinish = false
    c.mapJob = make([]Job, 0)
    c.reduceJob = make([]Job, 0)

    // 遍历输入文件，制作map任务清单
    for i, filename := range files {
        tempJob := Job{
            JobType: MapJob,
            OperateFileName: filename,
            JobId: i,
            JobState: UnFinish,
        }
        c.mapJob = append(c.mapJob, tempJob)
    }

    // 根据nRedcue, 制作reduce任务清单
    for i := 0; i < nReduce; i++ {
        tempJob := Job{
            JobType: ReduceJob,
            JobId: i,
            JobState: UnFinish,
        }
        c.reduceJob = append(c.reduceJob, tempJob)
    }
    
    // 将所有任务载入队列
    c.allJob = make(map[MRJobType][]Job)
    c.allJob[MapJob] = c.mapJob
    c.allJob[ReduceJob] = c.reduceJob


	c.server()
	return &c
}
