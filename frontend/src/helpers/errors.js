export const responseErrors = error => {
  if (error.response) {
    const { errors } = error.response.data
    return errors
  } else if (error.message) {
    return error.message
  } else if (error.request) {
    return error.request
  } else {
    return error.toString()
  }
}
