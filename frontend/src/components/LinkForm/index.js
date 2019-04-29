import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import {
  Form,
  FormGroup,
  Label,
  Input,
  FormFeedback,
  Alert,
  Row,
  Col,
  Button
} from 'reactstrap'

import * as linksActions from 'reducers/links/actions'

const mapStateToProps = state => ({
  form: state.links.form
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({ ...linksActions }, dispatch)
})

class LinkForm extends Component {
  static propTypes = {
    form: PropTypes.object.isRequired,
    actions: PropTypes.shape({
      createLink: PropTypes.func.isRequired,
      onLinkFieldChanged: PropTypes.func.isRequired,
      hideAlert: PropTypes.func.isRequired
    }).isRequired,
    onCancel: PropTypes.func.isRequired,
    onSuccess: PropTypes.func.isRequired
  }

  hideAlert = () => {
    this.props.actions.hideAlert()
  }

  onFieldChanged = (field, value) => {
    this.props.actions.onLinkFieldChanged(field, value)
  }

  onSubmit = (e) => {
    e.preventDefault()

    const { url } = this.props.form.fields

    const link = { url }

    this.props.actions.createLink(link)
  }

  render = () => {
    const { onCancel, form } = this.props
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
          <Label htmlFor="url">URL</Label>
          <Input type="url" name="url" id="url" placeholder="Enter the URL you want to shorten" onChange={(e) => this.onFieldChanged('url', e.target.value)} />
          <FormFeedback valid={!form.fields.urlHasError}>{form.fields.urlError}</FormFeedback>
        </FormGroup>
        <Row>
          <Col>
            <Button color="secondary" className="float-left" disabled={form.isFetching || !form.isValid} onClick={onCancel}>Cancel</Button>
            <Button color="primary" className="float-right" disabled={form.isFetching || !form.isValid}>Save</Button>
          </Col>
        </Row>
      </Form>
    )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(LinkForm)
