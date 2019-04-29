import {
  LOGIN,
  SIGNUP,
  LOGOUT
} from 'reducers/auth/constants'

export default state => {
  switch (state.state) {
    case LOGOUT:
      return state.setIn(['form', 'isValid'], true)
    case SIGNUP:
      if (state.form.fields.name !== '' &&
          state.form.fields.email !== '' &&
          state.form.fields.password !== '' &&
          state.form.fields.passwordConfirmation !== '' &&
          !state.form.fields.nameHasError &&
          !state.form.fields.emailHasError &&
          !state.form.fields.passwordHasError &&
          !state.form.fields.passwordConfirmationHasError) {
        return state.setIn(['form', 'isValid'], true)
      }
      return state.setIn(['form', 'isValid'], false)
    case LOGIN:
      if (state.form.fields.email !== '' &&
          state.form.fields.password !== '' &&
          !state.form.fields.emailHasError &&
          !state.form.fields.passwordHasError) {
        return state.setIn(['form', 'isValid'], true)
      }
      return state.setIn(['form', 'isValid'], false)
    default:
      return state
  }
}
