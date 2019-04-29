import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter, Route } from 'react-router-dom'

import Root from './Root'
import App from './containers/App'
import Login from './containers/Login'
import Signup from './containers/Signup'
import Home from './containers/Home'
import Links from './containers/Links'
import requireAuth from './hoc/requireAuth'

import './sass/app.scss'

ReactDOM.render(
  <Root>
    <BrowserRouter>
      <App>
        <Route path="/login" component={Login} />
        <Route path="/signup" component={Signup} />
        <Route path="/links" component={requireAuth(Links)} />
        <Route path="/" exact component={requireAuth(Home)} />
      </App>
    </BrowserRouter>
  </Root>,
  document.getElementById('root')
)
