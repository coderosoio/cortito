import { Record } from 'immutable'

const Form = Record({
  disabled: false,
  error: null,
  alertVisible: false,
  isValid: true,
  isFetching: false,
  fields: new (Record({
    url: '',
    urlHasError: false,
    urlError: ''
  }))()
})

const InitialState = Record({
  form: new Form(),
  links: []
})

export default InitialState
