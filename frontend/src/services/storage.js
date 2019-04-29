import store from 'store'

export default class StorageService {
  constructor () {
    this.store = store
  }

  set = (key, value) => (
    this.store.set(key, value)
  )

  get = (key, defaultValue = null) => {
    const value = this.store.get(key)
    if ((value === 'undefined') && (defaultValue !== null) && (defaultValue !== 'undefined')) {
      this.set(key, defaultValue)
      return defaultValue
    }
    return value
  }

  delete = key => (
    this.store.remove(key)
  )
}
