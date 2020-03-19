# TODO

```text
user
 | - request (涉及鉴权) with jwt
 | ---> [http] cas client 
         | - 询问是否有该用户 / jwt 是否有效
         | ---> [RPC] cas server
                  | - if 身份正确 ， cas client 缓存用户身份信息
```
```text
user
 | - request (注销)
 | ---> [http] cas server 
          | - 确认是否有该用户 / jwt 是否有效
          | ---> [RPC] cas clients (异步)
                  | - 修正缓存
  
```

注册登陆方式：

- email
- telephone
- oauth
  - github
  
```text
email/telephone

cas client (emailAddress) 
   -> cas server
       | - save email id and 验证码 答案
       -> send email to emailAddress
           | - user get  验证码 

cas client
   -> send 验证码答案 to cas server
```  
  
  
// cas 只做 rpc 服务
 
 
