import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import {
  Table
} from 'reactstrap'

import * as linksActions from 'reducers/links/actions'
import LinkItem from 'components/LinkItem'

const mapStateToProps = state => ({
  links: state.links.links
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({ ...linksActions }, dispatch)
})

class LinksList extends Component {
  static propTypes = {
    links: PropTypes.array.isRequired,
    actions: PropTypes.shape({
      listLinks: PropTypes.func.isRequired
    })
  }

  componentDidMount = () => {
    this.props.actions.listLinks()
  }

  render = () => {
    const { links } = this.props

    let linkItems = (
      <tr>
        <td colSpan="5">
          <div className="text-center">
            <h3>There are no links yet.</h3>
          </div>
        </td>
      </tr>
    )

    if (links.length > 0) {
      linkItems = links.map(link => (
        <LinkItem key={link.id} link={link} />
      ))
    }

    return (
      <Table>
        <thead>
          <tr>
            <th>Hash</th>
            <th>URL</th>
            <th>Visits</th>
            <th>Last visit</th>
            <th>Created</th>
          </tr>
        </thead>
        <tbody>
          {linkItems}
        </tbody>
      </Table>
    )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(LinksList)
