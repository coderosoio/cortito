import Api from 'services/api'

export default class LinksService extends Api {
  listLinks = () => (
    this.get('/shortener/links/')
  )

  createLink = (link) => (
    this.post('/shortener/links/', link)
  )

  shortenedLink = (hash) => (
    `${this.options.shortHost}/${hash}`
  )
}
