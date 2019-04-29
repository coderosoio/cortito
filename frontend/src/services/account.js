import Api from 'services/api'

export default class AccountService extends Api {
  signup = (user) => (
    this.post('/account/users/', user)
  )

  updateUser = user => (
    this.put(`/account/users/${user.id}`, user)
  )
}
