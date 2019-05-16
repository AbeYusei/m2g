# Easy CLI tool for converting .mov 2 .gif

## installation

you have to install ffmpeg CLI, and Go Version 1.11 or higher.

ffmpeg install example for your OSX local maschine.

```
brew install ffmpeg
```

install this CLI package.

```
go get -u github.com/AbeYusei/m2g/cmd/m2g
```

## how to use

```
m2g /file/to/movfile.mov
```
## run test

```
go test -v ./... -count=1
```
