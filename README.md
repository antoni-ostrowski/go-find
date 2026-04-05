# go-find

Simple `find` command clone.

## Development

```bash
just dev
```

## Build

```bash
just build
```

```bash
go-find -p <path> -q <query>
```

- `-p` - path to search (default: `.`)
- `-q` - query string to match in file paths (default: `""` matches all)

## Examples

```bash
go-find -p . -q ".go"
go-find -p ./src -q "test"
go-find -q "config.json"
go-find -p /home/user/projects -q "main"
```
