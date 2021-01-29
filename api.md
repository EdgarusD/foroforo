## GET

### Get all threads

#### `/thread`

**Request**

```json
{
    "offset": Integer,
    "limit": Integer,
}
```

**Response**

```json
{
    "threads": [
        {
            "id": String,
            "author": String,
            "title": String,
            "date": String,
        }
    ]
}
```

### Get thread

#### `/thread/:id`

**Response**

```json
{
    "id": String,
    "title": String,
    "body": String,
    "date": String,
    "comments": [
        {
            "id": String,
            "author": String,
            "comment": String,
            "date": String,
        }
    ],
}
```

### Get comment

#### `/thread/:id/comment/:id`

**Response**

```json
{
    "author": String,
    "comment": String,
    "date": String,
}
```

## POST

### Post thread

#### `/thread`

**Request**

```json
{
    "author": String,
    "title": String,
    "body": String,
}
```

**Response**

```json
{
    "id": String,
    "key": String,
}
```

### Post comment to thread

#### `/thread/:id/comment`

**Request**

```json
{
    "author": String,
    "comment": String,
}
```

**Response**

```json
{
    "id": String,
    "key": String,
}
```

## DELETE

### Delete thread

#### `/thread/:id`

**Request**

```json
{
    "key": String,
}
```

### Delete comment

#### `/thread/:id/comment/:id`

**Request**

```json
{
    "key": String,
}
```

## UPDATE

### Update thread

#### `/thread/:id`

**Request**

```json
{
    "key": String,
    "title": String | null,
    "body": String | null,
}
```

### Update comment

#### `/thread/:id/comment/:id`

**Request**

```json
{
    "key": String,
    "comment": String | null,
}
```
