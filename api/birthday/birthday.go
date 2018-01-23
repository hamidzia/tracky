// Package birhtday provides all the logic related to birthday records.
package birthday

import (
    "context"
	"time"
    "os"
    "encoding/csv"

    common "github.com/hamidzia/tracky/api/common"
    log "google.golang.org/appengine/log"
)

type Birthday struct {
    Name string
	Date time.Time
}
// Get all the birthdays from data file.
// path: comma separated CSV file (Full_Name,Date_Of_Birth)
func ReadBirthdays(c context.Context, path string)([]Birthday, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close() 

    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return nil, err
    }

    log.Infof(c, "Total Records: %d", len(lines))

    var birthdays []Birthday
    for _, line := range lines {
        t, err := time.Parse(common.DatePattern, line[1])
        if err != nil {
            log.Infof(c, "%v", err)
            continue
        }
        
        birthdays = append(birthdays, Birthday{Name: line[0], Date: t})
    }
    return birthdays, nil
}
