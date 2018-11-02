#go-package-plantuml

执行build

# 命令
pkgviewer 默认在当前目录生成uml.svg文件

# 修改
显示struct方法, 函数, type定义
自动生成svg文件, 需要java环境支撑

# 待优化
一个web页面,自动显示svg, 同时可以输入code目录, 以及ignore过滤条件

一种是分析包中的所有文件。

另一种是从入口分析整个依赖, 先在包级别，然后每个包具体方法调用与接口依赖

进一步，待考虑这是执行的trace，每一个方法的调用顺序




+ 根据demo
分析struct带包名
函数, 方法
调用关系
依赖关系

需要控制是否显示

+ 根据cellnet
调优
以及怎样排除不需要的包