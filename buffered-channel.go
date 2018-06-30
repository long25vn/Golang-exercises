 package main 
  
 import ( 
     "fmt" 
     "math/rand" 
     "sync" 
     "time" 
 ) 
  
 const ( 
     numWorkers = 4 
     numTasks = 10 
 ) 
  
 var wg sync.WaitGroup 
  
 func init() { 
     rand.Seed(time.Now().Unix()) 
 } 
  
 func main() { 
     // Tạo kênh đa chứa các công việc 
     tasks := make(chan string, numTasks) 
  
     wg.Add(numWorkers) 
     for gr := 1; gr <= numWorkers; gr++ { 
         go worker(tasks, gr) 
     } 
  
     // Tạo chuỗi các công việc cần thực thi cho vào kênh 
     for t := 1; t <= numTasks; t++ { 
         tasks <- fmt.Sprintf("công việc thứ %d.", t) 
     } 
  
     close(tasks) 
  
     wg.Wait() 
 } 
  
 // Thực hiện tác vụ, thực nhất là tạo và xử lý một goroutine 
 func worker(tasks chan string, worker int) { 
     defer wg.Done() 
  
     for { 
         // Chờ nhận công việc 
         task, ok := <-tasks 
         if !ok { 
             // Hết việc 
            fmt.Printf("Người thứ %d: Hoàn thành các công việc được giao!\n", worker) 
             return 
         } 
  
         fmt.Printf("Người thứ %d: Bắt đầu thực hiện %s\n", worker, task) 
  
         // Giả lập thời gian thực hiện công việc 
         sleep := rand.Int63n(100) 
         time.Sleep(time.Duration(sleep) * time.Millisecond) 
  
        fmt.Printf("Người thứ %d: Thực hiện xong %s\n", worker, task) 
    } 
 } 
