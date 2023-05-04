# logger

## Installation

```bash
go get -u github.com/metadiv-io/logger
```

## Setup log folder

Setup environment variable `LOG_FOLDER_PATH` to specify the log folder path.

## Highlights

* logger.Info(v ...any)

* logger.Error(v ...any)

* logger.Debug(v ...any)

* logger.Prefix("prefix").Info(v ...any)

* logger.Prefix("prefix").Error(v ...any)

* logger.Prefix("prefix").Debug(v ...any)
