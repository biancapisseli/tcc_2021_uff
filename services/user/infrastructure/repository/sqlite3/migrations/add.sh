
#!/bin/bash

MIGRATION_ID="$(date +"%Y%m%d%H%M%S")"
CAMEL_CASE_TITLE=$(echo "migration_$1" | sed -r 's/_([a-z])/\U\1/gi' | sed -r 's/^([A-Z])/\l\1/')
METHOD_TITLE="${CAMEL_CASE_TITLE}${MIGRATION_ID}"
METHOD_CALL="if err := m.${METHOD_TITLE}(); err != nil {\n\t\treturn err\n\t}"
CALLER_FILE="list.go"
NEW_MIGRATION_FILE="${MIGRATION_ID}_$1.go"

if [ -f "$NEW_MIGRATION_FILE" ]; then
    echo "migration exists"
    exit 1
fi

echo "
package migrations
import (
    \"fmt\"
)
func (m *migration) ${METHOD_TITLE}() error {
    version := \"${MIGRATION_ID}\"
    if _, ok := m.list[version]; ok {
        return fmt.Errorf(\"duplicated migration '%s_$1.go'\", version)
    }
    m.list[version] = []string{
        
    }
    return nil
}
" > $NEW_MIGRATION_FILE

sed -i "s/return nil/${METHOD_CALL}\n\treturn nil/" $CALLER_FILE