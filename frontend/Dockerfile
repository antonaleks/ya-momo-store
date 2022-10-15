### STAGE 1: Biuld Environment ###
FROM node:lts-alpine as builder
ENV VUE_APP_API_URL=http://localhost/rest/
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

### STAGE 2: Production Environment ###
FROM nginx:stable-alpine
COPY --from=builder /usr/src/app/dist /usr/share/nginx/html
RUN touch /usr/share/nginx/html/metrics.html
COPY .nginx/default.conf /etc/nginx/templates/default.conf.template
CMD ["nginx", "-g", "daemon off;"]
ENV API_HOST=localhost
ENV API_PORT=8081
EXPOSE 80