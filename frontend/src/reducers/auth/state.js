import { Record } from 'immutable'

import {
  SIGNUP
} from 'reducers/auth/constants'

const Form = Record({
  state: SIGNUP,
  disabled: false,
  error: null,
  alertVisible: false,
  isValid: true,
  isFetching: false,
  fields: new (Record({
    name: '',
    nameHasError: false,
    nameError: '',

    email: '',
    emailHasError: false,
    emailError: '',

    password: '',
    passwordHasError: false,
    passwordError: '',

    passwordConfirmation: '',
    passwordConfirmationHasError: false,
    passwordConfirmationError: '',

    showPassword: false
  }))()
})

const InitialState = Record({
  authenticated: false,
  form: new Form()
})

export default InitialState
