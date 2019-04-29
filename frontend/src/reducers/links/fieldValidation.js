import validate from 'validate.js'
import _ from 'underscore'

const constraints = {
  url: {
    url: {
      allowLocal: true
    }
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
