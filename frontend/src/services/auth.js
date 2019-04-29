import Api from 'services/api'

export default class AuthService extends Api {
  requestToken = (email, password) => (
    this.post('/account/auth/', { email, password })
  )

  revokeToken = () => (
    this.delete('/account/auth/')
  )

  getCurrentToken = () => (
    this.storage.get('token')
  )

  getCurrentUser = () => (
    this.storage.get('user')
  )

  setCurrentToken = (token) => (
    this.storage.set('token', token)
  )

  setCurrentUser = (user) => (
    this.storage.set('user', user)
  )
}
