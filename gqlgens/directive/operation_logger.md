# gql操作日志directive

## 使用方法
```text
1. 在query和mutation中加入指令 @operationLogger(objectType: ObjectType!,operType: MutationType) on FIELD_DEFINITION

2.  例子 common.graphql 中
type Query{
    t(req:treq):tresp @operationLogger(objectType: idea)
}

type Mutation{
    t(req:treq):tresp @operationLogger(objectType: idea,operType: update)
}

3. 可扩充
enum ObjectType{
    idea
    design
}
同时扩充 module/log/enum/objectType.go

4. 方法入参统一为 req

5. 
```