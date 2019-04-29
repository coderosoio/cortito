import { createBrowserHistory } from 'history'

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

  ON_AUTH_FIELD_CHANGED,

  TOGGLE_SHOW_PASSWORD,

  CLEAR_AUTH_ERRORS,

  HIDE_ALERT,

  LOGIN,
  LOGOUT,
  SIGNUP
} from 'reducers/auth/constants'
import { responseErrors } from 'helpers'
import AuthService from 'services/auth'
import AccountService from 'services/account'

const history = createBrowserHistory()
const authService = new AuthService()
const accountService = new AccountService()

export const hideAlert = () => ({
  type: HIDE_ALERT
})

export const onAuthFieldChanged = (field, value) => ({
  type: ON_AUTH_FIELD_CHANGED,
  payload: { field, value }
})

export const toggleShowPassword = () => ({
  type: TOGGLE_SHOW_PASSWORD
})

export const clearAuthErrors = () => ({
  type: CLEAR_AUTH_ERRORS
})

export const loginState = () => ({
  type: LOGIN
})

export const logoutState = () => ({
  type: LOGOUT
})

export const signupState = () => ({
  type: SIGNUP
})

export const loginRequest = () => ({
  type: LOGIN_REQUEST
})

export const loginSuccess = (token, user) => ({
  type: LOGIN_SUCCESS,
  payload: { token, user }
})

export const loginFailure = errors => ({
  type: LOGIN_FAILURE,
  payload: errors
})

export const tokenRequest = () => ({
  type: TOKEN_REQUEST
})

export const tokenSuccess = (token, user) => ({
  type: TOKEN_SUCCESS,
  payload: { token, user }
})

export const tokenFailure = errors => ({
  type: TOKEN_FAILURE,
  payload: errors
})

export const logoutRequest = () => ({
  type: LOGOUT_REQUEST
})

export const logoutSuccess = () => ({
  type: LOGOUT_SUCCESS
})

export const logoutFailure = errors => ({
  type: LOGOUT_FAILURE,
  payload: errors
})

export const signupRequest = () => ({
  type: SIGNUP_REQUEST
})

export const signupSuccess = user => ({
  type: SIGNUP_SUCCESS,
  payload: { user }
})

export const signupFailure = errors => ({
  type: SIGNUP_FAILURE,
  payload: errors
})

export const getToken = () => dispatch => {
  dispatch(tokenRequest())
  const token = authService.getCurrentToken()
  if (!token) {
    dispatch(tokenFailure('You are not logged in'))
    const location = history.location.pathname
    if (location !== '/signup' || location !== '/login') {
      history.replace('/login')
    }
  } else {
    const user = authService.getCurrentUser()
    dispatch(tokenSuccess(token, user))
  }
}

export const signup = user => dispatch => {
  dispatch(signupRequest())
  return accountService.signup(user)
    .then(response => {
      const user = response.data

      dispatch(clearAuthErrors())
      dispatch(signupSuccess(user))

      history.push('/login')
    })
    .catch(error => {
      const errors = responseErrors(error)
      dispatch(signupFailure(errors))
    })
}

export const login = (email, password) => dispatch => {
  dispatch(loginRequest())
  return authService.requestToken(email, password)
    .then(response => {
      const { user, token } = response.data

      authService.setCurrentToken(token)
      authService.setCurrentUser(user)

      dispatch(clearAuthErrors())
      dispatch(loginSuccess(token, user))

      history.replace('/')
    })
    .catch(error => {
      const errors = responseErrors(error)
      dispatch(loginFailure(errors))
    })
}

export const logout = () => dispatch => {
  dispatch(loginRequest())
  return authService.revokeToken()
    .then(() => {
      dispatch(logoutSuccess())
    })
    .catch(error => {
      const errors = responseErrors(error)
      dispatch(logoutFailure(errors))
    })
    .then(() => {
      authService.setCurrentToken(null)
      authService.setCurrentUser(null)

      history.replace('/')
    })
}
