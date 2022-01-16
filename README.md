## product name - cts price checker


### what is this?
a simple cli tool to get prices if cyptocurrencies and tokens against fiat currency

### why am i making it?
to get familiar with the language syntax. this is my first project since starting the Tour of Go

![gif of working prototype](https://media3.giphy.com/media/2vogqNpFQywWC0z0j5/giphy.gif)
<center style=color:grey;>gif of working prototype</center>

### want to use it?
you'll need to get an API key from nomics.com 

or any crypto price API tbh and just modify the `CurrencyResponse`struct in `model/pcmodel.go` to unmarshall the json data it recieves properly.

```
go run server.go
```
or 
```
generate an executable and run that
```

### how it works
makes a request to the nomics API, and prints the related price metrics on the CLI

#### technologies

- Go
- ❤️

### what's coming next
nothing at all. good first project. i learnt to make requests with `net/http`, got to use the Go bootleg Classes and Object alternatives, and played around with pointers. nothing more with this but on to better things

### want to help make this better?
text a crypto bro, tell them they're `ngmi`

[//]: # (So depending on use case, you may want to keep the documentation short. in that case you may not need to include the last two sections or you can)
