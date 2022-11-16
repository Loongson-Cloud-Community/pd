1. 该项目本身无架构相关的代码，主要是依赖的库文件包含架构相关代码   

2. 具体修改查看：git show 71cb52c586c000724c7dea1d583278962987b728   

3. 构建  
 
（1）构建环境
> cr.loongnix.cn/library/golang:1.19-alpine    

（2）构建二进制
> make   

（3）构建镜像   
> docker build -t xxx .   

备注：源码中go目录下的文件是适配LA架构的代码，在使用时将其替换为GOPATH目录下的go文件夹


