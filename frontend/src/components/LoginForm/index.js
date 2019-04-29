import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import {
  Row,
  Col,
  Form,
  FormGroup,
  Label,
  Input,
  Button,
  FormFeedback,
  Alert
} from 'reactstrap'
import { Link } from 'react-router-dom'

import * as authActions from 'reducers/auth/actions'

const mapStateToProps = state => ({
  form: state.auth.form
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({ ...authActions }, dispatch)
})

class LoginForm extends Component {
  static propTypes = {
    actions: PropTypes.shape({
      login: PropTypes.func.isRequired,
      toggleShowPassword: PropTypes.func.isRequired,
      onAuthFieldChanged: PropTypes.func.isRequired,
      hideAlert: PropTypes.func.isRequired
    }),
    form: PropTypes.object.isRequired
  }

  onSubmit = (e) => {
    e.preventDefault()

    const { email, password } = this.props.form.fields

    this.props.actions.login(email, password)
  }

  onFieldChanged = (field, value) => {
    this.props.actions.onAuthFieldChanged(field, value)
  }

  toggleShowPassword = () => {
    this.props.actions.toggleShowPassword()
  }

  hideAlert = () => {
    this.props.actions.hideAlert()
  }

  render = () => {
    const { form } = this.props
    let formErrors
    if (form.error) {
      formErrors = (
        <Alert color="danger" visible={form.error} isOpen={form.alertVisible} toggle={this.hideAlert}>
          {form.error || ''}
        </Alert>
      )
    }
    return (
      <Form onSubmit={this.onSubmit}>
        {formErrors}
        <FormGroup>
          <Label htmlFor="email">Email address</Label>
          <Input type="email" name="email" id="email" placeholder="Enter your email address" onChange={(e) => this.onFieldChanged('email', e.target.value)} invalid={form.fields.emailHasError} />
          <FormFeedback valid={!form.fields.emailHasError}>{form.fields.emailError}</FormFeedback>
        </FormGroup>
        <FormGroup>
          <Label htmlFor="password">Password</Label>
          <Input type={ form.fields.showPassword ? 'text' : 'password' } name="password" id="password" placeholder="Enter your password" onChange={(e) => this.onFieldChanged('password', e.target.value)} invalid={form.fields.passwordHasError} />
          <FormFeedback valid={!form.fields.passwordHasError}>{form.fields.passwordError}</FormFeedback>
        </FormGroup>
        <FormGroup check>
          <Label check>
            <Input type="checkbox" onChange={this.toggleShowPassword} /> Show password
          </Label>
        </FormGroup>
        <Row>
          <Col>
            <Button color="primary" className="float-right" disabled={form.isFetching || !form.isValid}>Log in</Button>
          </Col>
        </Row>
        <Link to="/signup" replace className="float-right">
          Don't have an account? Sign up here!
        </Link>
      </Form>
    )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(LoginForm)
