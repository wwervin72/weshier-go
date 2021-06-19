FROM nginx:1.17-alpine

ADD ./dist /usr/share/nginx/html
ADD ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
