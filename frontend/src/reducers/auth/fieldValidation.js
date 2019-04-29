import validate from 'validate.js'
import _ from 'underscore'

const constraints = {
  name: {
    presence: true
  },
  email: {
    presence: true,
    email: true
  },
  password: {
    presence: true,
    length: {
      minimum: 6,
      maximum: 128
    }
  },
  passwordConfirmation: {
    presence: true,
    equality: 'password'
  }
}

export default (state, action) => {
  const { field } = action.payload
  const subject = state.form.fields
  const result = validate(subject, constraints)
  const errors = result && result[field] ? result[field].join('\n') : undefined
  const hasError = !_.isUndefined(errors)
  const hasErrorKey = `${field}HasError`
  const errorKey = `${field}Error`

  return state.setIn(['form', 'fields', hasErrorKey], hasError)
              .setIn(['form', 'fields', errorKey], errors)
              .setIn(['form', 'isValid'], !hasError)
}
