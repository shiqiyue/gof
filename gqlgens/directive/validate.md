# gql参数验证directive

## 使用方法
```text
1、在graphql文件中加入指令的定义
directive @validate(rules: String!) on INPUT_FIELD_DEFINITION

2、在需要校验参数的地方使用指令
input ReportPageReq{
    Code: String
    CurrentPage: Int! @validate(rules: "min=1")
    PageSize: Int!  @validate(rules: "min=1,max=500")

}

3、使用gqlgen生成代码后，将Validate函数配置到Config
r := &Resolver{}
c := Config{
		Resolvers: r,
	}
c.Directives.Validate = directive.Validate
handler.NewDefaultServer(NewExecutableSchema(c))
```
