# behavior3go
forked from https://github.com/magicsea/behavior3go

golang behavior tree,from https://github.com/behavior3

## 简介

## 示例 Examples

- [load_from_tree](https://github.com/dyc92/behavior3go/tree/master/examples/load_from_tree)  ：从导出的树文件加载示例
- [load_from_project](https://github.com/dyc92/behavior3go/tree/master/examples/load_from_project) ：从导出的工程文件加载示例
- [load_from_rawproject](https://github.com/dyc92/behavior3go/tree/master/examples/load_from_rawproject) ：从原生工程文件加载示例
- [subtree](https://github.com/dyc92/behavior3go/tree/master/examples/subtree) ：子树的使用示例（需要[专用的编辑器分支版本](https://github.com/dyc92/behavior3editor)）
- [mmoarpg](examples/mmoarpg/zt.b3)：一个mmoarpg类型游戏的行为示例，可使用桌面版编辑器打开可以查看(Projects->open project)

## 完整示例

[io类游戏示例](https://github.com/dyc92/h5game/tree/master/server)  
bin/b3.json为行为树的数据，在编辑器中新建任意工程，选择（Project->Import->Tree as json）导入树就可以还原工程，如图。

## 更新

* 添加子树支持 SubTree 节点，需要编辑器修改node导出category字段

## 其他的参考

- 一些节点的概念 http://note.youdao.com/noteshare?id=4f46dc2144ea62b55f597630f5e666b4&sub=FF0B4E1D7916473E8DFC7242CFC62B69
- 其他项目基础知识 https://www.behaviac.com/concepts/
- 其他人写的behavior3go代码介绍 https://blog.csdn.net/u013272009/article/details/77131226
- behavior3的erlang实现，同时里面有一些行为树基础分享：https://github.com/dong50252409/behavior3erl
## FAQ
- Q:子树的相同记忆节点的黑板信息是重复的？
```
A:是个问题，但是由于下一次进入会onopen重置节点黑板信息，并不会有错误表现。如果遇到BUG表现欢迎提出。
```
- Q:Tick里的target如何调用
```
A:用在ai里，一般target就这个ai的拥有者，拥有者有blackboard的成员。
```
- Q:如何设计打断一个进行中的状态？   
A:参考https://github.com/dyc92/behavior3go/issues/15
## TODO
- [ ] 参数类型化
- [ ] 参数支持传递黑板值利用格式“@变量名”
- [ ] 子树支持自定义参数传递
## 上线项目

* [丛林大作战](https://www.taptap.com/app/31608)

## 联系

QQ群:285728047
