import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import * as authActions from 'reducers/auth/actions'
import Header from 'components/Header'

const mapStateToProps = state => ({
  authenticated: state.auth.authenticated
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({ ...authActions }, dispatch)
})

class App extends Component {
  static propTypes = {
    children: PropTypes.oneOfType([
      PropTypes.arrayOf(PropTypes.node),
      PropTypes.node
    ]),
    actions: PropTypes.objectOf(PropTypes.func),
    authenticated: PropTypes.bool.isRequired,
    user: PropTypes.object
  }

  componentDidMount = () => {
    this.props.actions.getToken()
  }

  render = () => {
    const { children } = this.props
    return (
      <div>
        <Header />
        <main>
          {children}
        </main>
      </div>
    )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(App)
