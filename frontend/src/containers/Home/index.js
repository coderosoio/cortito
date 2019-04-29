import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import {
  Container,
  Row,
  Col,
  Button,
  Modal,
  ModalHeader,
  ModalBody
} from 'reactstrap'

import * as linksActions from 'reducers/links/actions'
import LinksList from 'components/LinksList'
import LinkForm from 'components/LinkForm'

const mapStateToProps = state => ({
  authenticated: state.auth.authenticated,
  links: state.links.links
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({ ...linksActions }, dispatch)
})

class Home extends Component {
  static propTypes = {
    authenticated: PropTypes.bool.isRequired,
  }

  state = {
    showLinkForm: false
  }

  toggleLinkForm = () => {
    this.setState({
      showLinkForm: !this.state.showLinkForm
    })
  }

  render = () => {
    const { authenticated, links } = this.props
    let linksList
    if (authenticated) {
      linksList = <LinksList links={links} />
    }
    return (
      <Container>
        <Row>
          <Col><h3>Home</h3></Col>
        </Row>
        <Row>
          <Col sm="12" md={{ size: 4, offset: 8 }}>
            <Button color="success" onClick={this.toggleLinkForm} className="float-right">
              New short URL
            </Button>
            <Modal isOpen={this.state.showLinkForm} toggle={this.toggleLinkForm}>
              <ModalHeader toggle={this.toggleLinkForm}>New link</ModalHeader>
              <ModalBody>
                <LinkForm onSuccess={this.toggleLinkForm} onCancel={this.toggleLinkForm} />
              </ModalBody>
            </Modal>
          </Col>
        </Row>
        <Row>
          <Col>
            {linksList}
          </Col>
        </Row>
      </Container>
    )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Home)
