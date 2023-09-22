### 开闭原则
一个软件实体（如类、模块和函数）应当对扩展开放，对修改关闭。<br>
即在修改需求时不更改已有代码而是通过扩展来变化。
### 依赖倒置原则（Dependence Inversion Principle）
程序的设计要依赖于抽象接口，而不是方法的具体实现。
1. 高层业务逻辑：参考抽象层的接口来设计
2. 抽象层：抽象层的接口就是业务层可以使用的方法，业务层通过多态的方式利用接口指针来调用具体的实现方法
3. 实现层：根据抽象层来实现每个实现层的模块