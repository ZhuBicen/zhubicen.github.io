---
title: "如何遍历 multimap 中所有的 Key"
date: 2006-10-19T21:50:15+08:00
draft: false
tags: [C++]
---
此文为当时自己在网上(百度知道, 捂脸)的一个提问

```C++
typedef std::multimap<int ,std::string> MyMap;
typedef std::set<int> MySet;
void main()
{
MyMap map;
MySet set;
map.insert(std::make_pair(1,"12889898989"));
map.insert(std::make_pair(2,"343434323232"));
map.insert(std::make_pair(1,"2323232323"));
map.insert(std::make_pair(2,"2344423232323"));
｝
```

请问我如何才能做出一个循环，打印出`multimap`中的键为`1`和`2`呢

 google 以下, 此为答案:
```
int main()
{
  multimap<int,string> mm;
  mm.insert(make_pair(1, "a"));
  mm.insert(make_pair(1, "lemon"));
  mm.insert(make_pair(2, "peacock"));
  mm.insert(make_pair(3, "angel"));

  for( auto it = mm.begin(), end = mm.end(); it != end;
  		it = mm.upper_bound(it->first)) {
  			cout << it->first << ' ' << it->second << endl;     
       }
  
  return 0;
}
```

`p = c.upper_bound(k)` 的作用是: 
```
p points to the first element with key > k or c.end(), ordered container only
```

multimap 就是 `ordered` 的, 所以 `upper_bound` 可以逐个的跳过重复的 key

C++ 中还有 `unordered_multimap`, 这种情况下就不能使用 `upper_bound` 了, 只能一个个去过滤了

C++ 中的 `set` 或 `map` 可以结合不同的属性, `plain/unordered, plain/multi`,  `plain` 可以省略, 因此默认是 `ordered, plain` 属性

因此 C++ 中有这么多的 `set|map`

```C++
set   multiset  unordered_set unordered_multiset
map   multimap  unordered_map unordered_multimap
```

