### STAGE 1: Biuld Environment ###
FROM node:lts-alpine as builder
ARG VUE_APP_API_URL
ARG VUE_APP_VERSION
ENV VUE_APP_API_URL $VUE_APP_API_URL
ENV VUE_APP_VERSION $VUE_APP_VERSION
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

### STAGE 2: Production Environment ###
FROM nginx:stable-alpine
RUN touch /usr/share/nginx/html/check.html
COPY --from=builder /usr/src/app/dist /usr/share/nginx/html
COPY .nginx/default.conf /etc/nginx/conf.d/default.conf
CMD ["nginx", "-g", "daemon off;"]
EXPOSE 80