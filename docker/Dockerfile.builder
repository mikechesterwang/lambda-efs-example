FROM public.ecr.aws/amazonlinux/amazonlinux:latest
RUN yum install golang -y
WORKDIR /app
COPY . .
RUN mkdir -p /gocache
ENV GOCACHE /gocache
CMD [ "sh", "-c", "cd src && go mod tidy && go build -o ./out/handler main.go" ]