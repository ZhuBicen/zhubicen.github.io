---
title: "React Hooks"
date: 2020-05-28T11:06:17+08:00
draft: true
tags: []
---

在这篇文章中，我们简单实践一下 react hooks，这个新特性从 React 16.8 中引入。

如下的代码实现了，一个文本框，让用户输入一个字符串并且输入回车后，存如了 `SessionStorage`

```JavaScript
import React, { Component } from 'react';
import PropTypes from "prop-types"

export default class TokenForm extends React.Component {
    static propTypes = {
        setToken: PropTypes.func.isRequired
    };

    constructor(props) {
        super(props);
        this.state = { value: '' }
    }

    handleSubmit = event => {
        event.preventDefault();
        const { setToken } = this.props;
        const token = this.state.value;
        if (token) {
            setToken(token);
        }
    }
    handleChange = (event) => {
        this.setState({ value: event.target.value });
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit} onChange={this.handleChange}>
                <input type="text" value={this.state.value}></input>
            </form>
        )
    }
}
```

以下的代码，仅仅显示用户输入的字符串，并在适当的时候保存用户输入的字符串：

```JavaScript
import React, { Component } from 'react';
import './App.css';
import TokenForm from './TokenForm'
export default class App extends Component {

  state = {
    token: null
  };

  componentDidMount() {
    this.setState({ token: sessionStorage.getItem("token") });
  }

  setToken = token => {
    sessionStorage.setItem("token", token);
    this.setState({ token });
  }

  render() {
    const { token } = this.state;
    return (
      <div>
        <h1>Starrry Eyed</h1>
        {token ? token : <TokenForm setToken={this.setToken} />}
      </div>
    )
  }

};
```

我们可以使用 React Hook 来重构上面的 App 类，首先引入了 `useState`, `useState` 会返回一个状态，以及一个更改状态的函数，另外我们使用`useEffect` 来定义`setToken` 发生时的动作，此时把字符串保存到 `SessionStorage`
```JavaScript
import React, { useState, useEffect } from 'react';
import './App.css';
import TokenForm from './TokenForm'

export default function App() {

  const [token, setToken]  = useState(sessionStorage.getItem("token"));

  useEffect(()=> {
    sessionStorage.setItem("token", token);
  })

  return (
    <div>
      <h1>Starrry Eyed</h1>
      {token ? token : <TokenForm setToken={setToken} />}
    </div>
  )

}
```

effect 函数，在每次 Render 运行结束后运行

代码： https://github.com/ZhuBicen/my-app-react-hook 

More about effect:  https://reactjs.org/docs/hooks-effect.html 



