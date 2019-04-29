import axios from 'axios'
import merge from 'deepmerge'
import humps from 'humps'

import StorageService from 'services/storage'

export default class Api {
  constructor (options = {}) {
    if (new.target === Api) {
      throw new TypeError('Can not construct API instances directly')
    }
    this.options = Object.assign(this.options || {}, {
      apiHost: process.env.REACT_APP_API_HOST.replace(/\/$/, ''),
      host: process.env.REACT_APP_HOST.replace(/\/$/, ''),
      shortHost: process.env.REACT_APP_SHORT_HOST.replace(/\/$/, '')
    }, options)
    this.storage = new StorageService()
  }

  get = (url, params = {}) => (
    this._request({
      method: 'GET',
      url,
      params
    })
  )

  put = (url, data = {}, params = {}) => (
    this._request({
      method: 'PUT',
      url,
      data,
      params
    })
  )

  post = (url, data = {}, params = {}) => (
    this._request({
      method: 'POST',
      url,
      data,
      params
    })
  )

  delete = (url, data = {}, params = {}) => (
    this._request({
      method: 'DELETE',
      url,
      data,
      params
    })
  )

  _request = config => {
    config.data = humps.decamelizeKeys(config.data)
    config.params = humps.decamelizeKeys(config.params)

    const token = this.storage.get('token', '')
    const requestConfig = merge({
      baseURL: this.options.apiHost,
      url: config.url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': token && token.length > 0 ? `JWT ${token}` : ''
      },
      responseType: 'json',
      method: 'GET',
      params: {},
      transformRequest: data => {
        if (data) {
          data = humps.decamelizeKeys(data)
          data = JSON.stringify(data)
        }
        return data
      },
      transformResponse: data => {
        if (data) {
          data = humps.camelizeKeys(data)
        }
        return data
      }
    }, config)

    return axios(requestConfig)
  }
}
