---
title: "Spring Security With Token"
date: 2020-03-23T21:06:06+08:00
draft: true
tags: []
---


对于标准的 Web 应用来说 Spring Security 的默认行为使很方便使用的。它使用基于 cookie 的认证和会话（session）。另外，它还会自动地处理好 CSRF 令牌（避免中间人攻击）。在大部分情况下，你只需要为不同的 route 设定相应的授权，从相应的数据库中读取用户信息即可。


另一方面，如果你只想创建一些 REST API 为外部的服务（SPA 或者移动应用）使用，你可能不需要 基于 Session 的安全。这正使 JWT（JSON Web Token) 的使用场景。所有的信息都存储在 token 中，基于此就可以创建无需 Session 的服务（Session-Less Server)。


JWT 需要在每一个 HTTP 请求中携带，因此服务器可以使用 token 中携带的信息认证用户。至与如何携带 token可以有不同的选择。比如 使用 URL 参数 或者 Authorization header:

```
Authorization: Bearer <token string>
```

JWT 主要包括三个部分：

1. 头部 - 主要包括 token 的类型以及使用的 hash 算法
2. 载荷- 主要包括关于用户的信息
3. 签名- 签名用来验证 token 没有被更改过


