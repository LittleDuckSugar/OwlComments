# golang 1.17 used for this project
FROM golang:1.17
 
RUN mkdir /owlcomment
 
COPY ./API /owlcomment
 
WORKDIR /owlcomment

RUN go build -o server . 

CMD [ "/owlcomment/server" ]