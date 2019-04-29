import {
  LINKS_REQUEST,
  LINKS_SUCCESS,
  LINKS_FAILURE,

  CREATE_LINK_REQUEST,
  CREATE_LINK_SUCCESS,
  CREATE_LINK_FAILURE,

  ON_LINK_FIELD_CHANGED,

  HIDE_ALERT
} from "reducers/links/constants"
import { responseErrors } from 'helpers'
import LinksService from 'services/links'

const linksService = new LinksService()

export const hideAlert = () => ({
  type: HIDE_ALERT
})

export const onLinkFieldChanged = (field, value) => ({
  type: ON_LINK_FIELD_CHANGED,
  payload: { field, value }
})

export const linksRequest = () => ({
  type: LINKS_REQUEST
})

export const linksSuccess = links => ({
  type: LINKS_SUCCESS,
  payload: { links }
})

export const linksFailure = errors => ({
  type: LINKS_FAILURE,
  payload: errors
})

export const createLinkRequest = () => ({
  type: CREATE_LINK_REQUEST
})

export const createLinkSuccess = link => ({
  type: CREATE_LINK_SUCCESS,
  payload: { link }
})

export const createLinkFailure = errors => ({
  type: CREATE_LINK_FAILURE,
  payload: errors
})

export const listLinks = () => dispatch => {
  dispatch(linksRequest())
  return linksService.listLinks()
    .then(response => {
      const links = response.data

      dispatch(linksSuccess(links))
    })
    .catch(error => {
      const errors = responseErrors(error)

      dispatch(linksFailure(errors))
    })
}

export const createLink = link => dispatch => {
  dispatch(createLinkRequest())
  return linksService.createLink(link)
    .then(response => {
      const link = response.data

      dispatch(createLinkSuccess(link))
    })
    .catch(error => {
      const errors = responseErrors(error)

      dispatch(createLinkFailure(errors))
    })
}
