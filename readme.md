# goflames

Web App that generates the famous childhood game F L A M E.


### Pre-requisite:

1. Install [yarn](https://yarnpkg.com)/[npm](https://www.npmjs.com)



### Installation:
```sh
$ cd goflames
$ go mod tidy # get go packages
$ cd public
$ npm install # if using npm
$ yarn add # if using yarn. Do not combine!!

$ yarn run build # to build the css assets
$ npm run build # if using npm
```

### Building the Assets:
```sh
$ cd public
$ yarn build # if using yarn;
$ npm run build # if using npm
```

### Running:
```sh
$ cd <to/root/project>
$ cp -a .env.example .env
$ go run main.go
```

### Running with module reload:
```sh
$ go get -u github.com/cosmtrek/air # install air using Go

# without go, with linux
$ curl -fLo air https://git.io/linux_air

# copy .air.example.conf to .air.conf
$ cp -a .air.example.conf .air.conf

# run application using `air`
$ air
```


## Related:
[flame](https://github.com/rbo13/flame)

