import React from 'react'
import PropTypes from 'prop-types'
import { Provider } from 'react-redux'
import { createStore, applyMiddleware, compose } from 'redux'
import promiseMiddleware from 'redux-promise'
import thunk from 'redux-thunk'
import { createLogger } from 'redux-logger'

import reducers from './reducers'
import AuthInitialState from './reducers/auth/state'
import GlobalInitialState from './reducers/global/state'
import LinksInitialState from './reducers/links/state'

const getInitialState = () => ({
  auth: new AuthInitialState(),
  global: new GlobalInitialState(),
  links: new LinksInitialState()
})
const logger = createLogger()
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose

const Root = ({ children, initialState = getInitialState() }) => {
  const store = createStore(
    reducers,
    initialState,
    composeEnhancers(
      applyMiddleware(
        thunk,
        promiseMiddleware,
        logger
      )
    )
  )

  return (
    <Provider store={store}>
      {children}
    </Provider>
  )
}

Root.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]).isRequired,
  initialState: PropTypes.object
}

export default Root
