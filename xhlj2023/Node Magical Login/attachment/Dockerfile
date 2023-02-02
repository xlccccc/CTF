FROM node:lts-alpine3.15

WORKDIR /usr/src/app

COPY ./app ./

RUN npm --registry https://registry.npm.taobao.org install

EXPOSE 80

CMD ["./start.sh"]
