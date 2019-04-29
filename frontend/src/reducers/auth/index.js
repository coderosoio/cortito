import {
  TOKEN_REQUEST,
  TOKEN_SUCCESS,
  TOKEN_FAILURE,

  LOGIN_REQUEST,
  LOGIN_SUCCESS,
  LOGIN_FAILURE,

  LOGOUT_REQUEST,
  LOGOUT_SUCCESS,
  LOGOUT_FAILURE,

  SIGNUP_REQUEST,
  SIGNUP_SUCCESS,
  SIGNUP_FAILURE,

  LOGIN,
  SIGNUP,
  LOGOUT,

  TOGGLE_SHOW_PASSWORD,

  CLEAR_AUTH_ERRORS,

  HIDE_ALERT,

  ON_AUTH_FIELD_CHANGED
} from 'reducers/auth/constants'
import InitialState from 'reducers/auth/state'
import formValidation from 'reducers/auth/formValidation'
import fieldValidation from 'reducers/auth/fieldValidation'

const initialState = new InitialState()

export default (state = initialState, action) => {
  if (!(state instanceof InitialState)) return initialState.mergeDeep(state)

  switch (action.type) {
    case LOGOUT:
    case CLEAR_AUTH_ERRORS:
      return state.setIn(['form', 'state'], action.type)
        .setIn(['form', 'error'], null)
        .setIn(['form', 'isValid'], true)
        .setIn(['form', 'alertVisible'], false)
        .setIn(['form', 'isFetching'], false)
        .setIn(['form', 'fields', 'name'], '')
        .setIn(['form', 'fields', 'nameHasError'], false)
        .setIn(['form', 'fields', 'nameError'], '')
        .setIn(['form', 'fields', 'email'], '')
        .setIn(['form', 'fields', 'emailHasError'], false)
        .setIn(['form', 'fields', 'emailError'], '')
        .setIn(['form', 'fields', 'password'], '')
        .setIn(['form', 'fields', 'passwordHasError'], false)
        .setIn(['form', 'fields', 'passwordError'], '')
        .setIn(['form', 'fields', 'passwordConfirmation'], '')
        .setIn(['form', 'fields', 'passwordConfirmationHasError'], false)
        .setIn(['form', 'fields', 'passwordConfirmationError'], '')
    case LOGIN:
    case SIGNUP:
      return state.setIn(['form', 'state'], action.type)
        .setIn(['form', 'error'], null)
        .setIn(['form', 'disabled'], false)
        .setIn(['form', 'isFetching'], false)
        .setIn(['form', 'isValid'], true)
        .setIn(['form', 'alertVisible'], false)
    case SIGNUP_SUCCESS:
      return state.setIn(['form', 'isFetching'], false)
        .setIn(['form', 'isValid'], true)
        .setIn(['form', 'alertVisible'], false)
        .setIn(['form', 'error'], null)
        .setIn(['form', 'disabled'], false)
        .set('authenticated', false)
    case LOGIN_SUCCESS:
    case TOKEN_SUCCESS:
      return state.setIn(['form', 'isFetching'], false)
        .setIn(['form', 'isValid'], true)
        .setIn(['form', 'alertVisible'], false)
        .setIn(['form', 'error'], null)
        .setIn(['form', 'disabled'], false)
        .set('authenticated', true)
    case TOKEN_REQUEST:
    case LOGIN_REQUEST:
    case SIGNUP_REQUEST:
    case LOGOUT_REQUEST:
      return state.setIn(['form', 'isFetching'], true)
        .setIn(['form', 'disabled'], true)
        .setIn(['form', 'error'], null)
        .setIn(['form', 'alertVisible'], false)
    case TOKEN_FAILURE:
    case LOGIN_FAILURE:
    case SIGNUP_FAILURE:
    case LOGOUT_FAILURE:
      return state.setIn(['form', 'isFetching'], false)
        .setIn(['form', 'alertVisible'], action.payload !== null)
        .setIn(['form', 'disabled'], false)
        .setIn(['form', 'error'], action.payload)
        .set('authenticated', false)
    case LOGOUT_SUCCESS:
      return state.setIn(['form', 'isFetching'], false)
        .setIn(['form', 'isValid'], true)
        .setIn(['form', 'alertVisible'], false)
        .setIn(['form', 'error'], null)
        .setIn(['form', 'disabled'], false)
        .set('authenticated', false)
    case TOGGLE_SHOW_PASSWORD:
      return state.setIn(['form', 'fields', 'showPassword'], !state.form.fields.showPassword)
    case ON_AUTH_FIELD_CHANGED: {
      const { field, value } = action.payload
      let nextState = state.setIn(['form', 'fields', field], value)
        .setIn(['form', 'error'], null)
      return formValidation(fieldValidation(nextState, action), action)
    }
    case HIDE_ALERT:
      return state.setIn(['form', 'alertVisible'], false)
    default:
      return state
  }
}
