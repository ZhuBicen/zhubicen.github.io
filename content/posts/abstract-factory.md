---
title: "Abstract Factory"
date: 2009-06-04T20:23:51+08:00
draft: false
tags: [C++]
---

AF是一种对象创建型模式。

对象创建型模式：将实例化委托给另一个对象。类创建型模式用继承来改变被初始化的类。

在设计模式中提到：“AbstractFactory类通常用FactoryMethod实现，但也可以用Prototype实现。” 

AF模式侧重于产品的系列性，增加对某一产品的支持将比较困难。ＦＭ模式则没有此限制，ＦＭ又被称为虚构造函数。

http://code.google.com/p/tcplex/source/browse/trunk/Design_Pattern/af.cpp
```C++
#include <iostream>
using namespace std;

class MapSite
{
    virtual void Enter() = 0;
};
class Room:public MapSite
{
    virtual void Enter(){
        cout << "Entering room!" << endl;
    }
};
class Door:public MapSite
{
    virtual void Enter(){
        cout << "Entering door!" << endl;
    }
};
class Maze
{
public:
    Maze();
    void AddRoom(Room*){}
    void AddDoor(Door*){}
};
class MazeFactory
{
public:
    virtual Maze* MakeMaze()const{
        return new Maze();
    }
    virtual Room* MakeRoom()const{
        return new Room();
    }
    virtual Door* MakeDoor()const{
        return new Door();
    }
};
class MazeGame
{
public:
    Maze* CreateMaze(MazeFactory& factory){
        Maze* mz = factory.MakeMaze();
        Room* rm = factory.MakeRoom();
        Door* dr = factory.MakeDoor();
        mz->AddRoom(rm);
        mz->AddDoor(dr);
        return mz;
    }
};
```

http://code.google.com/p/tcplex/source/browse/trunk/Design_Pattern/af2.cpp

```C++
#include <iostream>
using namespace std;
class Clonable
{
public:
    virtual Clonable* Clone() = 0;
};

    
class MapSite
{
public:
    virtual void Enter() = 0;
};
class Room:public MapSite, Clonable
{
public:
    virtual void Enter(){
        cout << "Entering room!" << endl;
    }
    virtual Room* Clone(){
        return new Room();
    }
    
};
class Door:public MapSite, Clonable
{
public:
    virtual void Enter(){
        cout << "Entering door!" << endl;
    }
    virtual Door* Clone(){
        return new Door();
    }
};
class Maze:public Clonable
{
public:
    Maze();
    void AddRoom(Room*){}
    void AddDoor(Door*){}
    virtual Maze* Clone(){
        return new Maze();
    }    
};
class MazeFactory
{
public:
    MazeFactory(Maze* maze, Room* room, Door* door)
        :maze_(maze), room_(room), door_(door){}    
    
    virtual Maze* MakeMaze()const{
        return maze_->Clone();
    }
    virtual Room* MakeRoom()const{
        return room_->Clone();
    }
    virtual Door* MakeDoor()const{
        return door_->Clone();
    }
private:
    Maze* maze_;
    Door* door_;
    Room* room_;
};
class MazeGame
{
public:
    Maze* CreateMaze(MazeFactory& factory){
        Maze* mz = factory.MakeMaze();
        Room* rm = factory.MakeRoom();
        Door* dr = factory.MakeDoor();
        mz->AddRoom(rm);
        mz->AddDoor(dr);
        return mz;
    }
};

```