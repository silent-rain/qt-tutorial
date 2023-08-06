# 发出信号与接收信号
- 有时需要知道信号是由哪个控件发出的。对此Qt5提供了sender()方法。
- 我们创建了两个按钮。
- 我们通过调用sender()方法来判断信号源， 并将其名称显示在窗体的状态栏中。
- 注意: `当前linux输出问空`



## qt编译
```
go mod vendor

GOWORK=off qtdeploy -fast test desktop
```

## 效果展示

