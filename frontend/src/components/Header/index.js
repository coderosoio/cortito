import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import {
  Nav,
  Navbar,
  NavbarBrand,
  NavbarToggler,
  Collapse
} from 'reactstrap'

import * as authActions from 'reducers/auth/actions'
import * as globalActions from 'reducers/global/actions'

const mapStateToProps = state => ({
  authenticated: state.auth.authenticated,
  isOpen: state.global.navbarOpen
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    ...authActions,
    ...globalActions
  }, dispatch)
})

class Header extends Component {
  static propTypes = {
    actions: PropTypes.shape({
      toggleNavbar: PropTypes.func.isRequired
    })
  }

  toggleNavbar = () => {
    this.props.actions.toggleNavbar()
  }

  render = () => {
    const { isOpen } = this.props
    return (
      <header>
        <Navbar color="dark" dark expand="md">
          <NavbarBrand href="/">Cortito</NavbarBrand>
          <NavbarToggler onClick={this.toggleNavbar} />
          <Collapse isOpen={isOpen} navbar>
            <Nav className="ml-auto" navbar>

            </Nav>
          </Collapse>
        </Navbar>
      </header>
    )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Header)
