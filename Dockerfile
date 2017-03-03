FROM ubuntu
RUN git clone https://github.com/paisantosh/smartclient.git
RUN cd smartclient
RUN go build smartclient.go
EXPOSE 8000
CMD ["smartclient"]
