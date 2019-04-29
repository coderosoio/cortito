import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { withRouter } from 'react-router-dom'

import AuthService from 'services/auth'

const authService = new AuthService()

export default ChildComponent => {
  class RequireAuth extends Component {
    static propTypes = {
      history: PropTypes.shape({
        push: PropTypes.func.isRequired
      }).isRequired
    }

    componentDidMount = () => {
      this.ensureAuthenticated()
    }

    componentDidUpdate = () => {
      this.ensureAuthenticated()
    }

    ensureAuthenticated = () => {
      const token = authService.getCurrentToken()
      if (!token) {
        this.props.history.push('/login')
      }
    }

    render = () => {
      return (
        <ChildComponent {...this.props} />
      )
    }
  }

  return withRouter(RequireAuth)
}
