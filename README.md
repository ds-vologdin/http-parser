# Usage

```bash
./http-parser -h
Usage of ./http-parser:
  -max-worker int
        max of worker (default 5) (default 5)
  -word string
        word for count (default 'Go') (default "Go")
```

## Example

```bash
echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org/doc/' | ./http-parser -max-worker 10 -word Go
```