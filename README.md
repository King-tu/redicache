# redicache
用Golang实现的一个线程安全的类 redis的缓存。
	1、使用Golang的map实现Redicache键值对的添加(Set)、查询(Get)方法
	2、实现键的查询(keys)、删除(del)功能，其中keys命令可以匹配正则表达式
	3、利用map的key无序且唯一的特点，实现Redicache集合(set)元素的添加(Sadd)、删除(Srem)、遍历(Smembers)、是否存在(Sismember)等方法以及多个集合之间的交集(Sinter)和并集(Sunion)运算
	4、增加读写锁，实现线程安全地访问缓存数据
	5、实现Redicache缓存数据的持久化：定时保存缓存数据至磁盘文件，Redicache退出时自动存盘，初始化时自动加载磁盘文件的数据到内存中
	6、实现Redicache并发服务器(RediServer)、客户端(RediClient)，基于TCP通信协议。
