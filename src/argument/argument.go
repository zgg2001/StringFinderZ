package argument

type Argument struct {
    Path    string  //路径
    Find    string  //搜索内容
    Replace string  //替换内容
    Number  bool    //显示行号
}

var Arg = &Argument{}