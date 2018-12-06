
### Stocks

## Requirements

It's necessary Docker so as to run this project.

## Clone
```bash
    git clone https://github.com/joaquinicolas/go-react-stocks.git
```

## Run the app
```bash
    docker build -t iextrading .
    docker run -p 8080:8080 -d iextrading
```

Or

```bash
    docker run -p 8080:8080 -d joaquinnicolas96/iextrading
```

Now, Open your browser at http://localhost:8080/