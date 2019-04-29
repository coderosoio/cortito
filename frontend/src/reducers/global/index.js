import {
  TOGGLE_NAVBAR
} from "reducers/global/constants"
import InitialState from 'reducers/global/state'

const initialState = new InitialState()

export default (state = initialState, action) => {
  if (!(state instanceof InitialState)) return initialState.mergeDeep(state)

  switch (action.type) {
    case TOGGLE_NAVBAR:
      return state.set('navbarOpen', !state.navbarOpen)
    default:
      return state
  }
}
