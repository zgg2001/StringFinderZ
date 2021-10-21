package argument

type Argument struct {
    Path    string  //路径
    Find    string  //搜索内容
    Replace string  //替换内容
    IsFind  bool    //是否为找到模式
    Number  bool    //显示行号
    AllYes  bool    //是否全部许可替换
    Quick   bool    //是否快速替换
}

var Arg = &Argument{}
