## gorm-pagination

---

#### Usage

---
```shell
go get github.com/dogukanayd/gorm-pagination
```

Just create the Config struct from the package and use Paginate method

```go
	p, nor := (&paginator.Config{
		Page:    1,
		PerPage: 2,
		Path:    c.Path(),
		Sort:    "id desc",
	}).Paginate(db, &cb)

    return c.JSON(http.StatusOK, p)
```
Example Response:
```json
{
  "current_page": 1,
  "data": [
    {
      "id": 3,
      "user_name": "John"
    },
    {
      "id": 2,
      "campaign_name": "Jane"
    },
    {
      "id": 1,
      "campaign_name": "Luke"
    }
  ],
  "first_page_url": "127.0.0.1:3030/users/get?page=1&per_page=3",
  "from": 0,
  "last_page": 748,
  "last_page_url": "127.0.0.1:3030/users/get?page=748&per_page=3",
  "next_page_url": "127.0.0.1:3030/users/get?page=2&per_page=3",
  "path": "/users/get",
  "per_page": 3,
  "prev_page_url": "127.0.0.1:3030/users/get?page=1&per_page=3",
  "to": 0,
  "total": 2245
}
```
For unit tests it's provide interface;
```go
// Paginator ..
type Paginator interface {
	Paginate(db *gorm.DB) (Result, *gorm.DB)
}
```
