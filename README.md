# bufferManager.go
`bufferManager` demonstrates a technique to decrease pressure on GC in Go.

Due to [overhead of type assertion](https://groups.google.com/d/msg/golang-nuts/lKHCCFC9WfU/v6bggyb7HrMJ) and [lack of generics](http://golang.org/doc/faq#generics), there is currently not an efficient way to implement this into a general purpose library. If you wanna use it, just replace bufferManager.Data with your own data structure.

To run the test,
```
$ git clone git://github.com/songgao/bufferManager.go.git
$ cd bufferManager.go
$ chmod +x run_test.sh
$ ./run_test.sh
```

```
===========================
Running the test......

With bufferManager:                    768.742ms
Without buffermanager (Relying on GC): 6.497013s
===========================


===========================
pprof - with bufferManager:

Total: 64 samples
      26  40.6%  40.6%       26  40.6% reflect.Value.MapIndex
      13  20.3%  60.9%       61  95.3% runtime.chansend
      10  15.6%  76.6%       30  46.9% runtime.chanrecv
       5   7.8%  84.4%       21  32.8% runtime.unlock
       4   6.2%  90.6%       16  25.0% runtime.lock
       3   4.7%  95.3%        3   4.7% reflect.chancap
       2   3.1%  98.4%       34  53.1% runtime.chanrecv1
       1   1.6% 100.0%       27  42.2% runtime.chansend1
       0   0.0% 100.0%       61  95.3% reflect.Value.Index
===========================


===========================
pprof - without bufferManager:

Total: 640 samples
     615  96.1%  96.1%      640 100.0% reflect.Value.MapIndex
      24   3.8%  99.8%       24   3.8% runtime.FixAlloc_Free
       1   0.2% 100.0%        3   0.5% reflect.Value.call
       0   0.0% 100.0%      640 100.0% reflect.Value.Index
       0   0.0% 100.0%        3   0.5% runtime.MCache_Alloc
       0   0.0% 100.0%        3   0.5% runtime.MCentral_AllocList
       0   0.0% 100.0%        2   0.3% runtime.MGetSizeClassInfo
       0   0.0% 100.0%      640 100.0% runtime.chansend
       0   0.0% 100.0%      637  99.5% runtime.gc
===========================
```
