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
    user_type             : fk -> UserType             : required
```


### UserType

TODO: Intermediate table between subject to indicate skill level

```
    id   : int (PK) : required
    type : varchar  : required
```

### TutorSubject

```
    id      : int (PK)   : required
    user    : fk -> User : required
    subject : 1->1       : required
    rating  : 1-5        : required
```

### Calendar
```
    id      : int (PK)   : required
    user    : fk -> User : required
```

### Authentication Method

```
    id   : int (PK) : required
    type : varchar  : required
```

### Tutorship

```
    id             : int (PK)         : required
    tutor          : fk -> User       : required
    tutee          : fk -> User       : required
    written_review : text
    rating         : 1-5
    homework       : fk 1->m Homework
    classes        : fk 1->m Class
    subject        : m->m Subject     : required
```

### Messages

```
    id        : int (PK)   : required
    receiver  : fk -> User : required
    sender    : fk -> User : required
    time_sent : datetime   : required
    content   : text       : required
```

### Homework

```
    id      : int (PK) : required
    content : text     : required
```

### Reservation
Reservation is potential class in calendar???

```
    id                     : int (PK)   : required
    class                  : 1->1 Class
    calendar *->1 Calendar
    time                   : datetime   : required
    approved               : bool       : required
```

### Class

```
    id          : int (PK) : required
    description : text     : required
    rating      : 1-5      :
```
