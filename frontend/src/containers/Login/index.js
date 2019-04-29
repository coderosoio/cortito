import React, { Component } from 'react'
import {
  Container,
  Row,
  Col
} from 'reactstrap'

import LoginForm from "components/LoginForm"

export default class Login extends Component {
  render = () => {
    return (
      <Container>
        <Row>
          <Col sm="" md={{ size: 4, offset: 4 }}>
            <h3>Log in</h3>
            <LoginForm />
          </Col>
        </Row>
      </Container>
    )
  }
}
