>分布式缓存模拟器：
- 可作为NameNode，DataNode，Client进行启动，在不同节点间建立连接与通讯。

![image-20221114233404707](https://cdn.jsdelivr.net/gh/Gavinyjb/images@master/images/202211142334612.png)

![image-20221114233442876](https://cdn.jsdelivr.net/gh/Gavinyjb/images@master/images/202211142334943.png)

###命令列表
####NameNode

```shell
go run start.go
```

####DataNode
```shell
go run start.go -nodeType DataNode -myPort 3004 -myName slave4
```


####Client
```shell
go run start.go -nodeType Client  -command getDNList
```