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
      signup: PropTypes.func.isRequired,
      toggleShowPassword: PropTypes.func.isRequired,
      onAuthFieldChanged: PropTypes.func.isRequired,
      hideAlert: PropTypes.func.isRequired
    }),
    form: PropTypes.object.isRequired
  }

  onSubmit = (e) => {
    e.preventDefault()

    const { name, email, password, passwordConfirmation } = this.props.form.fields

    const user = { name, email, password, passwordConfirmation }

    this.props.actions.signup(user)
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
          <Label htmlFor="name">Your name</Label>
          <Input id="name" type="text" name="name" placeholder="Enter your name" onChange={(e) => this.onFieldChanged('name', e.target.value)} invalid={form.fields.nameHasError} />
          <FormFeedback valid={!form.fields.nameHasError}>{form.fields.nameError}</FormFeedback>
        </FormGroup>
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
        <FormGroup>
          <Label htmlFor="passwordConfirmation">Confirm your password</Label>
          <Input type={ form.fields.showPassword ? 'text' : 'password' } name="passwordConfirmation" id="passwordConfirmation" placeholder="Please confirm your password" onChange={(e) => this.onFieldChanged('passwordConfirmation', e.target.value)} invalid={form.fields.passwordConfirmationHasError} />
          <FormFeedback valid={!form.fields.passwordConfirmationHasError}>{form.fields.passwordConfirmationError}</FormFeedback>
        </FormGroup>
        <FormGroup check>
          <Label check>
            <Input type="checkbox" onChange={this.toggleShowPassword} /> Show password
          </Label>
        </FormGroup>
        <Row>
          <Col>
            <Button color="primary" className="float-right" disabled={form.isFetching || !form.isValid}>Sign up</Button>
          </Col>
        </Row>
        <Link to="/login" replace className="float-right">
          Already have an account? Log in here!
        </Link>
      </Form>
    )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(LoginForm)
