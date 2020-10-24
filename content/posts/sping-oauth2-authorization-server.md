---
title: "创建并使用 PostMan 测试 OAuth2 授权服务器"
date: 2020-10-19T22:38:45+08:00
draft: false
tags: [Spring, OAuth2]
---


在分布式系统中，授权系统是必不可少的，作为最流行的分布式系统架构之一的 Spring Cloud 提供了对 OAuth2 授权服务器的支持。

Spring Cloud 对与 OAuth2 的支持都集成在 `spring.security.oatuh2` 包中，使用 Spring 开发 OAuth2 授权服务器，只需做三件事：

##  使用 `@EnableAuthorizationServer`
```Java
@RestController
@EnableAuthorizationServer
@SpringBootApplication
public class OauthServerApplication {
    public static void main(String[] args) {
       	SpringApplication.run(OauthServerApplication.class, args);
    }
}
```

## 注册 Client 信息
通过 继承`AuthorizationServerConfigurerAdapter` 注册 `Client` 信息，要注册的信息包括：Client ID，Client  Secret，授权模式，重定向 URL 以及权限范围。

其中 client ID 和 Client Secret 用与 Client Authentication

```Java
@Configuration
public class Oauth2Config extends AuthorizationServerConfigurerAdapter {
    //...
    @Override
    public void configure(ClientDetailsServiceConfigurer clients) throws Exception 	   {
        clients.inMemory()
                .withClient("eagleeye")
                .secret("{noop}thisissecret")
                .authorizedGrantTypes("authorization_code")
                .redirectUris("http://localhost:9090/callback")
                .scopes("webclient", "mobileclient");
    }
    //...
}
```
## 配置 Resource Owner  的信息
在配置了 Client 的信息后，另外只需配置普通的 Authentication 用户信息就可以了

```Java
    @Override
    protected void configure(AuthenticationManagerBuilder auth) throws Exception {
        auth.inMemoryAuthentication()
            .withUser("user").password("{noop}password").roles("USER")
            .and()
            .withUser("admin").password("{noop}password").roles("USER", "ADMIN");
    }
```

## 提供 Rest 接口返回用户信息
```Java
	@RequestMapping(value = {"/user"}, produces="application/json")
	public Map<String, Object> user(OAuth2Authentication user) {
		Map<String, Object> userInfo = new HashMap<>();
		userInfo.put(
				"user",
				user.getUserAuthentication().getPrincipal());
		userInfo.put(
				"authorities",
				AuthorityUtils.authorityListToSet(user.getUserAuthentication().getAuthorities()));
		return userInfo;
	}
```

## 测试 OAuth2 授权服务器

使用 Post Man 测试，创建一个 Request，在 `Authorization` 字段中选择 `OAuth 2.0`：

![image-20201024223552641](sping-oauth2-authorization-server.assets/image-20201024223552641.png)

然后点击 `get access token` ，填入 OAuth2 相关的信息：

![image-20201024224954660](sping-oauth2-authorization-server.assets/image-20201024224954660.png)


点击 ` Request Token`， 会跳出 `Resource Owner` 的授权信息：

![image-20201024224519454](sping-oauth2-authorization-server.assets/image-20201024224519454.png)

输入用户名： “user”，用户名：“password”，点击 “Sign in”

就可以获取到 Access Token：

![image-20201024231829082](sping-oauth2-authorization-server.assets/image-20201024231829082.png)

获取到 Access Token 后，就可以通过 Rest 接口获取用户的信息。点击 Use Token 按钮，请求 `/user` 就可以获取用户信息：

![image-20201024231905085](sping-oauth2-authorization-server.assets/image-20201024231905085.png)

本文使用的相关代码在 [GitHub](!https://github.com/ZhuBicen/simple_oauth2/tree/master/oauth2-server)

另外附上 OAuth2 服务器相关的资源：

1. [一个 OAuth2 授权服务器的例子：说明文档](!https://howtodoinjava.com/spring-boot2/oauth2-auth-server/Source )，代码在 [github](!https://github.com/lokeshgupta1981/SpringExamples/tree/master/oauth2)
2. [这篇文章讲述 OAuth2 的概念](!https://medium.com/@rameez.s.shaikh/spring-boot-oauth2-authorization-code-grant-beb9b3b589f3)
3. [Spring Boot 提供的 OAuth2 起步教程](!https://spring.io/guides/tutorials/spring-boot-oauth2/)
4. [OAuth2 授权码模式实现](!https://segmentfault.com/a/1190000012275317，https://www.cnblogs.com/hellxz/p/oauth2_oauthcode_pattern.html)
5. [稍微有点老的文章](!https://raymondhlee.wordpress.com/2014/12/21/implementing-oauth2-with-spring-security/)
6. [Spring  Security 和 Gateway](! https://spring.io/blog/2019/08/16/securing-services-with-spring-cloud-gateway)
7. [有详细操作步骤的一个例子](!https://developer.okta.com/blog/2019/03/12/oauth2-spring-security-guide)

