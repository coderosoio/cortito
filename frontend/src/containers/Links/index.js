import React, { Component } from 'react'
import {
  Container,
  Row,
  Col,
  Button,
  Modal,
  ModalHeader,
  ModalBody
} from 'reactstrap'

import LinksList from 'components/LinksList'
import LinkForm from 'components/LinkForm'

export default class Links extends Component {
  state = {
    showLinkForm: false
  }

  toggleLinkForm = () => {
    this.setState({
      showLinkForm: !this.state.showLinkForm
    })
  }

  render = () => {
    return (
      <Container>
        <Row>
          <Col sm="12" md={{ size: 4, offset: 8 }}>
            <Button color="success" onClick={this.toggleLinkForm}>
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
            <LinksList />
          </Col>
        </Row>
      </Container>
    )
  }
}
