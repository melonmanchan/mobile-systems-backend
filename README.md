# Mobile systems backend

## Database model schema

### User

```
    id                    : int (PK)                   : required
    first_name            : varchar                    : required
    last_name             : varchar                    : required
    email                 : varchar                    : required
    location              : geolocation                : required
    authentication_method : fk -> AuthenticationMethod : required
    user_role             : fk -> UserRole             : required
```


### UserRole

```
    id   : int (PK) : required
    type : varchar  : required
```

### Authentication Method

```
    id   : int (PK) : required
    type : varchar  : required
```

### Tutorship

```
    id       : int (PK)         : required
    tutor    : fk -> User       : required
    tutee    : fk -> User       : required
    homework : fk 1->m Homework
    classes  : fk 1->m Class
```

### Messages

```
    id        : int (PK)   : required
    receiver  : fk -> User : required
    sender    : fk -> User : required
    time_sent : datetime   : required
    content   : text       : required
```
