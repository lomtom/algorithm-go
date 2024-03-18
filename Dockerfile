FROM httpd:2.4
COPY ./docs/dist /usr/local/apache2/htdocs/
EXPOSE 80