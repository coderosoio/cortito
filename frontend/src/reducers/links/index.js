import {
  LINKS_REQUEST,
  LINKS_SUCCESS,
  LINKS_FAILURE,

  CREATE_LINK_REQUEST,
  CREATE_LINK_SUCCESS,
  CREATE_LINK_FAILURE,

  ON_LINK_FIELD_CHANGED,

  HIDE_ALERT
} from "reducers/links/constants"
import formValidation from 'reducers/links/formValidation'
import fieldValidation from 'reducers/links/fieldValidation'
import InitialState from 'reducers/links/state'

const initialState = new InitialState()

export default (state = initialState, action) => {
  if (!(state instanceof InitialState)) initialState.mergeDeep(state)

  switch (action.type) {
    case LINKS_REQUEST:
    case CREATE_LINK_REQUEST:
      return state.setIn(['form', 'isFetching'], true)
        .setIn(['form', 'disabled'], true)
        .setIn(['form', 'alertVisible'], false)
        .setIn(['form', 'error'], null)
    case LINKS_SUCCESS: {
      const { links } = action.payload
      return state.setIn(['form', 'isFetching'], false)
        .setIn(['form', 'disabled'], false)
        .setIn(['form', 'alertVisible'], false)
        .setIn(['form', 'error'], null)
        .set('links', links)
    }
    case CREATE_LINK_SUCCESS: {
      const { link } = action.payload
      const links = [...state.links, link]
      return state.setIn(['form', 'isFetching'], false)
        .setIn(['form', 'disabled'], false)
        .setIn(['form', 'alertVisible'], false)
        .setIn(['form', 'error'], null)
        .setIn(['form', 'url'], '')
        .set('links', links)
    }
    case LINKS_FAILURE:
    case CREATE_LINK_FAILURE:
      return state.setIn(['form', 'isFetching'], false)
        .setIn(['form', 'disabled'], false)
        .setIn(['form', 'alertVisible'], true)
        .setIn(['form', 'error'], action.payload)
    case HIDE_ALERT:
      return state.setIn(['form', 'alertVisible'], false)
    case ON_LINK_FIELD_CHANGED: {
      const { field, value } = action.payload
      let nextState = state.setIn(['form', 'fields', field], value)
        .setIn(['form', 'error'], null)
      return formValidation(fieldValidation(nextState, action), action)
    }
    default:
      return state
  }
}
