import { combineReducers } from 'redux'

import auth from 'reducers/auth'
import global from 'reducers/global'
import links from 'reducers/links'

const rootReducer = combineReducers({
  auth,
  global,
  links
})

export default rootReducer
