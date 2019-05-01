FROM node:latest AS builder

COPY frontend /cortito/frontend

WORKDIR /cortito/frontend

ENV NODE_ENV=production

RUN yarn

RUN yarn build

FROM alpine:latest

COPY --from=builder /cortito/frontend/build /cortito/

