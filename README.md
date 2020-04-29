# Usage

```bash
./http-parser -h
Usage of ./http-parser:
  -max-workers int
        max count of workers (default 5)
  -word string
        word for count (default "Go")
```

## Example

```bash
echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org/doc/' | ./http-parser -max-workers 10 -word Go
```
