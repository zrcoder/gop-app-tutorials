# Gop app tutorials 
Tutorials of writing web apps in Go+ language.

## Prerequisite
Look at the `go.mod`/`gop.mod` files, version of Go and Go+ required are:

```
go 1.17
go+ 1.0
```

## How to run

```shell
cd ${tutorial dir}
make run ${port}
```

> For example:
>
> ```shell
> cd 01-HellowWorld
> make run port=9990
> ```
> and there will be a log like below:
> ```
> ...
> serving on [http://localhost:9990]
> ```
>
> Now you can visit http://localhost:9990 on the browser.