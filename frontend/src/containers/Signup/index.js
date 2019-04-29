import React, { Component } from 'react'
import {
  Container,
  Row,
  Col
} from 'reactstrap'

import SignupForm from 'components/SignupForm'

export default class Signup extends Component {
  render = () => {
    return (
      <Container>
        <Row>
          <Col sm="12" md={{ size: 6, offset: 3 }} lg={{ size: 4, offset: 4 }}>
            <h3>Sign up</h3>
            <SignupForm />
          </Col>
        </Row>
      </Container>
    )
  }
}
