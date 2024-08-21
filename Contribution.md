If you want to contribute to this project, start with an existing issue or propose a new one. 

# Development

## Build

```bash
VERSION=$(semver info v)
go build -ldflags "-X main.version=$VERSION" .   
```

# Promote

## Create a tag

