# Test Assignment

#### Requirements:
- [Docker](https://docs.docker.com/get-docker/)
- [Golang](https://golang.org/doc/install)

#### Usage:
To run server locally you can use either docker:
```bash
make start
```
or Go 
```bash
make dev
```

Use curl or whatever you want to test endpoint:
Example:
```bash
curl --location --request GET 'http://localhost:3000/v1/calculate?a=true&b=true&c=true&d=10.32&e=100&f=8&mode=base'
```

To run test:
```bash
make test
```

By default server runs on 3000 port, use SERVER_PORT env variable
to change it

#### Technical description:
For implementing http server I used Gorilla mux cause it
provides a very simple interface for managing routes and server.

The main implementation of this assignment is in switcher package.
The idea that there might be different switchers in my particular
case it is logical one. I created struct and put formulas inside
so we can easily adjust function costructor and for instance pass 
formulas as parameters. To apply custom sets I use an optional 
parameter mode that can accept one of three values: base, custom1 or 
custom2. 


