---
title: "Spring Boot 升级过程中遇到的一个问题"
date: 2020-03-31T13:38:06+08:00
draft: false
tags: [Java]
---

最近升级了项目中使用的 Spring Boot 版本，从 1.3.5 升级到了 2.2.5。升级过程中遇到了一个问题，原有的 Interceptor 不工作了，现象是 这样的：

```
org.springframework.web.servlet.resource.ResourceHttpRequestHandler cannot be cast to org.springframework.web.method.HandlerMethod
```

相关的代码如下，它实现了对访问用户的记录：

```java
public class ActionInterceptor implements HandlerInterceptor {
	private Logger logger = LoggerFactory.getLogger(getClass());
	private ThreadLocal<Long> startTime = new ThreadLocal<>();
	
	@Override
	public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object methodObj) throws Exception {

		HandlerMethod method = (HandlerMethod) methodObj;
		VisitorLog visitLog = method.getMethodAnnotation(VisitorLog.class);
		if (visitLog != null) {
			logVisitors(request, visitLog);
		}

		RoleProtecting roleProtecting = method.getMethodAnnotation(RoleProtecting.class);
    }
```



抛出异常的在这一句：`HandlerMethod method = (HandlerMethod) methodObj`，原因在于 `methodObj` 的类型为`ResourceHttpRequestHandler`。 这是 Spring Boot 2（Spring 5.x)  的一个改变：对于静态资源的请求也会转发到 `interceptor` 。这就是这个异常的原因



Fix 的办法就很简单了，在强制转换前，判断一下 instance 的类型即可：

```Java
		if (!(methodObj instanceof HandlerMethod)) {
			return true;
		}
		HandlerMethod method = (HandlerMethod) methodObj;
```

另外一个方法，就是在添加 interceptor 时，指定 interceptor 要处理的请求的范围，比如使用：`excludePathPatterns` 把静态资源的 URL 排除，比如：

```java
@Configuration
public class ControllerConfiguration implements WebMvcConfigurer {

    private final AdminOnlyInterceptor adminInterceptor;

    @Autowired
    public ControllerConfiguration(AdminInterceptor adminInterceptor) {
        this.adminInterceptor = adminInterceptor;
    }

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(adminInterceptor)
                .addPathPatterns("/rest-api-root/**"); // White list paths
                //.excludePathPatterns("/static-resource-root/**"); // Black list paths
    }
}
```



这个 Spring Boot 的变化还是需要引起注意，升级的过程中其实要处理的问题还是比较多的，我在第一次尝试升级的时候，直接就把这个 interceptor 直接注释掉了，后面静下来看，才发现根本原因



参考： https://stackoverflow.com/questions/53460730/spring-boot-2-0-intercepting-the-handlermethod-of-a-request 