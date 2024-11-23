# FROM mdsol/node16
FROM gplane/pnpm:9.12.1-node18

WORKDIR /web/

COPY . /web/

RUN pnpm install
# RUN npm install
# RUN pnpm run build

# RUN pnpm run serve

# FROM nginx:alpine
# LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

# COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
# COPY --from=0 /web/dist /usr/share/nginx/html
# RUN cat /etc/nginx/nginx.conf
# RUN cat /etc/nginx/conf.d/my.conf
# RUN ls -al /usr/share/nginx/html

# CMD [ "pnpm run serve" ]