package threadpool

import (
	"dingjing/queue"
	"library/go-logging"
	"sync"
)

type ThreadpoolFunc func(i interface{})

type CThreadpool struct {
	queue      *queue.Queue
	quesize    int
	condition  *sync.Cond
	threadNum  int
	exit       bool
	threadFunc ThreadpoolFunc
	log        *logging.Logger
}

func New() *CThreadpool {
	tp := new(CThreadpool)
	tp.queue = queue.New()
	tp.condition = sync.NewCond(&sync.Mutex{})
	tp.threadNum = 0
	tp.quesize = 200
	tp.exit = false
	tp.threadFunc = nil
	tp.log = logging.MustGetLogger("thread pool")

	return tp
}

/* 设置任务队列大小 */
func (tp *CThreadpool) SetQueueSize(num int) {
	if num > 0 {
		tp.quesize = num
		tp.log.Infof("任务队列容量为: %d", num)
	}
}

/* 设置线程数 */
func (tp *CThreadpool) SetThreadNum(num int) {
	if num > 0 {
		tp.threadNum = num
		tp.log.Infof("将开启 %d 个线程", num)
	}
}

/* 设置线程中的处理函数 */
func (tp *CThreadpool) SetFunc(f ThreadpoolFunc) {
	if nil != f {
		tp.threadFunc = f
	}
}

/* 添加任务 */
func (tp *CThreadpool) AddWork(d interface{}) {
	tp.condition.L.Lock()
	if tp.queue.Size() >= tp.quesize {
		tp.log.Infof("任务队列已满，等待...")
		tp.condition.Wait()
	}

	tp.queue.Push(d)
	tp.condition.Signal()
	tp.log.Infof("已添加新任务!!!")
	tp.condition.L.Unlock()
}

/* 退出线程池 */
func (tp *CThreadpool) Exit() {
	tp.log.Infof("线程池开始退出...")
	tp.condition.L.Lock()
	tp.exit = true
	tp.condition.Broadcast()
	tp.condition.L.Unlock()
}

/* 开始线程 */
func (tp *CThreadpool) Run() {
	group := sync.WaitGroup{}
	for i := 0; i < tp.threadNum; i++ {
		group.Add(1)
		go func() {




			tp.log.Infof("线程 %d 已准备就绪...", i)
			for !tp.exit {
				tp.condition.L.Lock()

				// 没有任务
				for tp.queue.Size() <= 0 {
					tp.condition.Wait()
				}

				if nil != tp.threadFunc {
					// 取值 条件变量
					data := tp.queue.Pop()
					tp.condition.Signal()
					tp.condition.L.Unlock()
					tp.threadFunc(data)
				} else {
					panic("没有设置任务函数...")
				}
			}
			group.Done()
			tp.log.Infof("线程 %d 已退出...", i)
		}()
	}
	group.Wait()
}
