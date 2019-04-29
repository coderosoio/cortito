FROM node:latest

COPY ./frontend/package.json /cortito/frontend/
COPY ./frontend/yarn.lock /cortito/frontend/

WORKDIR /cortito/frontend

RUN yarn

CMD ["yarn", "start"]
