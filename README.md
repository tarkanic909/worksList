# worksList

Simple CLI app written in go, for finding all works from authors of specific book

## Download

Executable binaries can be found in `bin` directory

## How to use

### Options

| Option  | Description                                                | Type     | Default    | Required? |
| ------- | ---------------------------------------------------------- | -------- | ---------- | --------- |
| `-isbn` | ISBN of the book                                           | `string` | 1617291781 | No        |
| `-sort` | Option for sorting by count of revision. Options: asc/desc | `string` | asc        | No        |

### Example

```bash
worksList -book="Lord of the rings" -sort=asc
```

### Help

```bash
worksList -h --help
```
