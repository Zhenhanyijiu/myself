env GOPATH=/D_QAwYqH3be/yangyanan/myself/myself:/D_QAwYqH3be/yangyanan/gopath/ go build -o libtest.so -buildmode=c-shared test.go//.so
env GOPATH=/D_QAwYqH3be/yangyanan/myself/myself:/D_QAwYqH3be/yangyanan/gopath/ go build -o libtest.a -buildmode=c-archive test.go//.a

go build 可以指定buildmode。分为了多种模式。具体模式如下。

模式	说明
当前go版本	1.10.3
archive	编译成二进制文件。一般是静态库文件。 xx.a
c-archive	编译成C归档文件。C可调用的静态库。xx.a。注意要编译成此类文件需要import C 并且要外部调用的函数要使用 “//export 函数名” 的方式在函数上方注释。否则函数默认不会被导出。
c-shared	编译成C共享库。同样需要 import “C” 和在函数上方注释 // export xxx
default	对于有main包的直接编译成可执行文件。没有main包的，编译成.a文件
exe	编译成window可执行程序
plugin	将main包和依赖的包一起编译成go plugin。非main包忽略。【类似C的共享库或静态库。插件式开发使用】
————————————————
版权声明：本文为CSDN博主「zouxinjiang」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/github_33719169/article/details/84826983
https://blog.csdn.net/github_33719169/article/details/84826983