# worksList

Simple CLI app written in go, for finding all works from authors of specific book

## Download

Executable binaries can be found in `bin` directory

## How to use

### Options

| Option      | Description                                                | Type     | Default     | Required? |
| ----------- | ---------------------------------------------------------- | -------- | ----------- | --------- |
| `-olid`     | Open Library ID of the book                                | `string` | OL32011221M | No        |
| `-author`   | Option for sorting by author name. Options: asc/desc       | `string` | asc         | No        |
| `-revision` | Option for sorting by count of revision. Options: asc/desc | `string` | asc         | No        |

### Example

```bash
worksList -olid="OL32011221M" -revision=asc -author=asc
```

### Help

```bash
worksList -h --help
```
