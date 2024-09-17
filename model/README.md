# Model Package

This package provides CRUD operations for models in a PocketBase application.

## Overview

The `model` package defines a `Crud` interface and provides functions to perform common database operations such as finding all records, finding a record by ID, and saving a record.

## Interface

```go
type Crud interface {
    models.Model
    GetUser() string
    Validate() error
}
```

## Functions

### FindAll

```go
func FindAll[C Crud](model Crud, dao *daos.Dao, authRecord *models.Record) ([]C, error)
```

Finds all records of a given model that belong to the authenticated user.

### FindById

```go
func FindById(model Crud, dao *daos.Dao, authRecord *models.Record, id string) error
```

Finds a record by its ID and ensures it belongs to the authenticated user.

### Save

```go
func Save(model Crud, dao *daos.Dao) error
```

Validates and saves a model to the database.

## Dependencies

- [PocketBase](https://github.com/pocketbase/pocketbase)
- [dbx](https://github.com/pocketbase/dbx)

## Usage

Import the package and use the provided functions to interact with your models.

```go
import (
    "github.com/yourusername/yourrepo/model"
    "github.com/pocketbase/pocketbase/daos"
    "github.com/pocketbase/pocketbase/models"
)
```
