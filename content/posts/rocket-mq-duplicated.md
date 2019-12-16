---
title: "Rocket MQ 的 Message Listener 收到重复的消息"
date: 2019-07-23T14:43:31+08:00
draft: false
tags: [Java]
---

## 问题

Rocket MQ 的 MessageListener 每次都会收到同样的消息的两个副本，两个消息的 Message Id 是一样的


## 解决办法

Message Id 是一样的说明是同一个消息，这是 Rocket MQ 的重试机制的现象。原因是在第一次处理此消息的时候，发生了异常（比如 NullPointerException），导致处理函数异常退出，因此没有返回 `CONSUME_SUCCESS`, 因此 Rocket MQ 会重新使用此消息回调消息处理函数。重试的次数应该是可以设定的。 

因此在消息回调函数使用 try catch 是有助于发现和调查此类问题的

```
 consumer.registerMessageListener(new MessageListenerConcurrently() {
            @Override
            public ConsumeConcurrentlyStatus consumeMessage(List<MessageExt> msgs,
                ConsumeConcurrentlyContext context) {
                try {
                    System.out.printf("%s Receive New Messages: %s %n", Thread.currentThread().getName(), msgs);
                    // do something
                    return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
                }
                catch(Exception e) {
                    // logout the exception
                }
            }
        });
```