import { capitalize } from 'humanize'

export const responseErrors = error => {
  // console.log(error)
  // console.log('response: ', error.response)
  // console.log('message: ', error.message)
  // console.log('request: ', error.request)
  if (error.response) {
    const { error: errors } = error.response.data
    if (typeof errors === 'object') {
      return Object.keys(errors).map(key => (
        capitalize(errors[key])
      )).join('\n')
    } else if (typeof errors === 'string') {
      return errors
    }
    return errors.toString()
  } else if (error.message) {
    return error.message
  } else if (error.request) {
    return error.request
  } else {
    return error.toString()
  }
}
