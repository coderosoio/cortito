export default (state, action) => {
  if (state.form.fields.url !== '' &&
      !state.form.fields.urlHasError) {
    return state.setIn(['form', 'isValid'], true)
  }
  return state.setIn(['form', 'isValid'], false)
}
