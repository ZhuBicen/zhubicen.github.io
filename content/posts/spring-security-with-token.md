---
title: "在 Spring Security 中使用 Token"
date: 2020-03-23T21:06:06+08:00
draft: false
tags: [Java, Spring]
---



翻译来自于[*blog*](https://jakublesko.com/spring-security-with-jwt/)

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



## Token 举例

放在 `authorization header` 中的 token 大概看起来是这个样子：

```
Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJzZWN1cmUtYXBpIiwiYXVkIjoic2VjdXJlLWFwcCIsInN1YiI6InVzZXIiLCJleHAiOjE1NDgyNDI1ODksInJvbCI6WyJST0xFX1VTRVIiXX0.GzUPUWStRofrWI9Ctfv2h-XofGZwcOog9swtuqg1vSkA8kDWLcY3InVgmct7rq4ZU3lxI6CGupNgSazypHoFOA

```

正如你看到的，token 中的三个部分 - 头部，载荷 和签名 是以逗号分割的。头部和载荷都是一 Base64 编码的 JSON 对象

头部：

```json
{
  "typ": "JWT",
  "alg": "HS512"
}
```

载荷：

```json
{
  "iss": "secure-api",
  "aud": "secure-app",
  "sub": "user",
  "exp": 1548242589,
  "rol": [
    "ROLE_USER"
  ]
}
```

## 应用举例

在接下来的例子中，我们将创建两个 Rest API， 一个API 是对所用用户开放，另一个私有的 API 仅仅对授权用户开放

我们使用 [start.spring.io](https://start.spring.io/) 来生成应用程序框架，同时选择使用`Spring Security` `Spring Web` 依赖。其他的依赖你可以依据需要来选择

![image-20200328120010808](spring-security-with-token.assets/image-20200328120010808.png)



JJWT 提供了Java 对JWT 的支持，因此我们需要吧如下的依赖添加到 pom.xml 中

```xml
<dependency>
    <groupId>io.jsonwebtoken</groupId>
    <artifactId>jjwt-api</artifactId>
    <version>0.10.5</version>
</dependency>
<dependency>
    <groupId>io.jsonwebtoken</groupId>
    <artifactId>jjwt-impl</artifactId>
    <version>0.10.5</version>
    <scope>runtime</scope>
</dependency>
<dependency>
    <groupId>io.jsonwebtoken</groupId>
    <artifactId>jjwt-jackson</artifactId>
    <version>0.10.5</version>
    <scope>runtime</scope>
</dependency>
```



## Controllers

在我们例子中的 Controller 将会尽可能的简单。它们将会简单的返回一个消息 或者 在用户没有授权的情况下仅仅返回 403 的错误码

``` Java
@RestController
@RequestMapping("/api/public")
public class PublicController {

    @GetMapping
    public String getMessage() {
        return "Hello from public API controller";
    }
}
@RestController
@RequestMapping("/api/private")
public class PrivateController {

    @GetMapping
    public String getMessage() {
        return "Hello from private API controller";
    }
}
```

## Filters

首先，我们将定义一些常量用于 JWT 的生成与验证

> 注意：最好不要把 JWT 的签名应编码到应用中（在我们的例子中我们忽略了这一点），你应该使用环境变量或者 .properties 文件。另外，签名也需要有合适的长度，举例来说，HS512 算法需要一个至少512 字节的签名



```Java
public final class SecurityConstants {

    public static final String AUTH_LOGIN_URL = "/api/authenticate";

    // Signing key for HS512 algorithm
    // You can use the page http://www.allkeysgenerator.com/ to generate all kinds of keys
    public static final String JWT_SECRET = "n2r5u8x/A%D*G-KaPdSgVkYp3s6v9y$B&E(H+MbQeThWmZq4t7w!z%C*F-J@NcRf";

    // JWT token defaults
    public static final String TOKEN_HEADER = "Authorization";
    public static final String TOKEN_PREFIX = "Bearer ";
    public static final String TOKEN_TYPE = "JWT";
    public static final String TOKEN_ISSUER = "secure-api";
    public static final String TOKEN_AUDIENCE = "secure-app";

    private SecurityConstants() {
        throw new IllegalStateException("Cannot create instance of static util class");
    }
}
```

第一个 filter 用来对用户鉴权（Authentication），它使用URL参数提供的用户名和密码调用`authentication manager` 来验证

如果用户名密码使正确的，FILTER 将会生成 JWT token 并且通过 `authorization header`

```Java
public class JwtAuthenticationFilter extends UsernamePasswordAuthenticationFilter {

    private final AuthenticationManager authenticationManager;

    public JwtAuthenticationFilter(AuthenticationManager authenticationManager) {
        this.authenticationManager = authenticationManager;

        setFilterProcessesUrl(SecurityConstants.AUTH_LOGIN_URL);
    }

    @Override
    public Authentication attemptAuthentication(HttpServletRequest request, HttpServletResponse response) {
        var username = request.getParameter("username");
        var password = request.getParameter("password");
        var authenticationToken = new UsernamePasswordAuthenticationToken(username, password);

        return authenticationManager.authenticate(authenticationToken);
    }

    @Override
    protected void successfulAuthentication(HttpServletRequest request, HttpServletResponse response,
                                            FilterChain filterChain, Authentication authentication) {
        var user = ((User) authentication.getPrincipal());

        var roles = user.getAuthorities()
            .stream()
            .map(GrantedAuthority::getAuthority)
            .collect(Collectors.toList());

        var signingKey = SecurityConstants.JWT_SECRET.getBytes();

        var token = Jwts.builder()
            .signWith(Keys.hmacShaKeyFor(signingKey), SignatureAlgorithm.HS512)
            .setHeaderParam("typ", SecurityConstants.TOKEN_TYPE)
            .setIssuer(SecurityConstants.TOKEN_ISSUER)
            .setAudience(SecurityConstants.TOKEN_AUDIENCE)
            .setSubject(user.getUsername())
            .setExpiration(new Date(System.currentTimeMillis() + 864000000))
            .claim("rol", roles)
            .compact();

        response.addHeader(SecurityConstants.TOKEN_HEADER, SecurityConstants.TOKEN_PREFIX + token);
    }
}
```

第二个 Filter 将会拦截所有的 HTTP 请求，并且检查 `Authorization header` 中的 token 是否正确。比如 token 是否过期或者签名是否正确

如果 token 是有效的 filter 将会把鉴权信息存入 Spring Security 上下文中

```Java
public class JwtAuthenticationFilter extends UsernamePasswordAuthenticationFilter {

    private final AuthenticationManager authenticationManager;

    public JwtAuthenticationFilter(AuthenticationManager authenticationManager) {
        this.authenticationManager = authenticationManager;

        setFilterProcessesUrl(SecurityConstants.AUTH_LOGIN_URL);
    }

    @Override
    public Authentication attemptAuthentication(HttpServletRequest request, HttpServletResponse response) {
        var username = request.getParameter("username");
        var password = request.getParameter("password");
        var authenticationToken = new UsernamePasswordAuthenticationToken(username, password);

        return authenticationManager.authenticate(authenticationToken);
    }

    @Override
    protected void successfulAuthentication(HttpServletRequest request, HttpServletResponse response,
                                            FilterChain filterChain, Authentication authentication) {
        var user = ((User) authentication.getPrincipal());

        var roles = user.getAuthorities()
            .stream()
            .map(GrantedAuthority::getAuthority)
            .collect(Collectors.toList());

        var signingKey = SecurityConstants.JWT_SECRET.getBytes();

        var token = Jwts.builder()
            .signWith(Keys.hmacShaKeyFor(signingKey), SignatureAlgorithm.HS512)
            .setHeaderParam("typ", SecurityConstants.TOKEN_TYPE)
            .setIssuer(SecurityConstants.TOKEN_ISSUER)
            .setAudience(SecurityConstants.TOKEN_AUDIENCE)
            .setSubject(user.getUsername())
            .setExpiration(new Date(System.currentTimeMillis() + 864000000))
            .claim("rol", roles)
            .compact();

        response.addHeader(SecurityConstants.TOKEN_HEADER, SecurityConstants.TOKEN_PREFIX + token);
    }
}
```

## 安全配置

最后我们需要对 Spring Security 本身进行配置。配置都很简单，我们仅仅需要设置很少的细节：

+ 密码编码：我们使用了 bcypt
+ CORS 的配置
+ Authentication Manager - 在本例子中我们使用了内存中的鉴权，但是在实际的开发中，我们可能需要使用 `UserDetailsService`
+ 设置什么端点使安全的（需要登录），什么端点使公开的
+ 在 Security 上下文中添加两个 Filter
+ 禁用 Session 的管理 - 我们不需要 Session 所以配置不需要生成 session cookies

```Java
@EnableWebSecurity
@EnableGlobalMethodSecurity(securedEnabled = true)
public class SecurityConfiguration extends WebSecurityConfigurerAdapter {

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http.cors().and()
            .csrf().disable()
            .authorizeRequests()
            .antMatchers("/api/public").permitAll()
            .anyRequest().authenticated()
            .and()
            .addFilter(new JwtAuthenticationFilter(authenticationManager()))
            .addFilter(new JwtAuthorizationFilter(authenticationManager()))
            .sessionManagement()
            .sessionCreationPolicy(SessionCreationPolicy.STATELESS);
    }

    @Override
    public void configure(AuthenticationManagerBuilder auth) throws Exception {
        auth.inMemoryAuthentication()
            .withUser("user")
            .password(passwordEncoder().encode("password"))
            .authorities("ROLE_USER");
    }

    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }

    @Bean
    public CorsConfigurationSource corsConfigurationSource() {
        final UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        source.registerCorsConfiguration("/**", new CorsConfiguration().applyPermitDefaultValues());

        return source;
    }
}
```



## 测试

请求公开的 API

```
POST http://localhost:8080/api/authenticate?username=user&password=password
HTTP/1.1 200 
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJzZWN1cmUtYXBpIiwiYXVkIjoic2VjdXJlLWFwcCIsInN1YiI6InVzZXIiLCJleHAiOjE1NDgyNDYwNzUsInJvbCI6WyJST0xFX1VTRVIiXX0.yhskhWyi-PgIluYY21rL0saAG92TfTVVVgVT1afWd_NnmOMg__2kK5lcna3lXzYI4-0qi9uGpI6Ul33-b9KTnA
X-Content-Type-Options: nosniff
X-XSS-Protection: 1; mode=block
Cache-Control: no-cache, no-store, max-age=0, must-revalidate
Pragma: no-cache
Expires: 0
X-Frame-Options: DENY
Content-Length: 0
Date: Sun, 13 Jan 2019 12:21:15 GMT

<Response body is empty>

Response code: 200; Time: 167ms; Content length: 0 bytes
```



使用上一步得到的 token 请求私有 API

```
GET http://localhost:8080/api/private
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJzZWN1cmUtYXBpIiwiYXVkIjoic2VjdXJlLWFwcCIsInN1YiI6InVzZXIiLCJleHAiOjE1NDgyNDI1ODksInJvbCI6WyJST0xFX1VTRVIiXX0.GzUPUWStRofrWI9Ctfv2h-XofGZwcOog9swtuqg1vSkA8kDWLcY3InVgmct7rq4ZU3lxI6CGupNgSazypHoFOA
HTTP/1.1 200 
X-Content-Type-Options: nosniff
X-XSS-Protection: 1; mode=block
Cache-Control: no-cache, no-store, max-age=0, must-revalidate
Pragma: no-cache
Expires: 0
X-Frame-Options: DENY
Content-Type: text/plain;charset=UTF-8
Content-Length: 33
Date: Sun, 13 Jan 2019 12:22:48 GMT

Hello from private API controller

Response code: 200; Time: 12ms; Content length: 33 bytes
```

直接请求私有API

此时你会得到 403，因为你没有使用 token 访问私有的端点

```
GET http://localhost:8080/api/private
HTTP/1.1 403 
X-Content-Type-Options: nosniff
X-XSS-Protection: 1; mode=block
Cache-Control: no-cache, no-store, max-age=0, must-revalidate
Pragma: no-cache
Expires: 0
X-Frame-Options: DENY
Content-Type: application/json;charset=UTF-8
Transfer-Encoding: chunked
Date: Sun, 13 Jan 2019 12:27:25 GMT

{
  "timestamp": "2019-01-13T12:27:25.020+0000",
  "status": 403,
  "error": "Forbidden",
  "message": "Access Denied",
  "path": "/api/private"
}

Response code: 403; Time: 28ms; Content length: 125 bytes
```



## 结论

本文的目标是演示如何在 Spring Security 中使用 JWT。这个例子可以使用在你的产品中。另外，在本文中我没有想更深入的探讨 token 的刷新，验证等，或许在后面的文章中将会探讨这些主题

本文涉及的所有的代码都存放在： [github 仓库](https://github.com/keysh/jwt-security)

