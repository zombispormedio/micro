# 👩🏼‍💻 micro

> Minimal docker image for static frontend (less than 8Mb)

## Install

```bash
docker pull zombispormedio/micro
```

## Usage

```docker
FROM zombispormedio/micro:1.0.0

COPY public /var/www/public

ENV PORT 8085

ENV STATIC_DIR /var/www/public

EXPOSE 8085

CMD ["./micro"]
```

## Built with

- [go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software..
- [fasthttp](https://github.com/valyala/fasthttp) - Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http

## License

MIT © [Xavier Serrano](https://zombispormedio.github.io)
