
package migrations
import (
    "fmt"
)
func (m *migration) migrationInitial20211023173018() error {
    version := "20211023173018"
    if _, ok := m.list[version]; ok {
        return fmt.Errorf("duplicated migration '%s_initial.go'", version)
    }
    m.list[version] = []string{
        
    }
    return nil
}

