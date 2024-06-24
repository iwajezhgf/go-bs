# go-bs

Bcypt service on Go.

## API

**Hash password:**
```
POST /hash
Content-Type: application/x-www-form-urlencoded

raw=<password>&cost=<cost>
```

**Verify password:**
```
POST /verify
Content-Type: application/x-www-form-urlencoded

raw=<password>&hash=<hash>
```

## Run

```
go-bs -addr 127.0.0.1:8843
```