import React, { Component } from 'react'
import PropTypes from 'prop-types'

import LinksService from 'services/links'

const linksService = new LinksService()

export default class LinkItem extends Component {
  static propTypes = {
    link: PropTypes.shape({
      id: PropTypes.number.isRequired,
      hash: PropTypes.string.isRequired,
      url: PropTypes.string.isRequired,
      visits: PropTypes.number,
      lastVisit: PropTypes.string,
      createdAt: PropTypes.string.isRequired
    }).isRequired
  }

  render = () => {
    const { link } = this.props
    return (
      <tr>
        <td><a href={linksService.shortenedLink(link.hash)} target="_blank" rel="noopener noreferrer">{link.hash}</a></td>
        <td><a href={link.url} target="_blank" rel="noopener noreferrer">{link.url}</a></td>
        <td>{link.visits || 0}</td>
        <td>{link.lastVisit || 'Not visited yet'}</td>
        <td>{link.createdAt}</td>
      </tr>
    )
  }
}
