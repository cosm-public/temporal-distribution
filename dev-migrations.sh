#!/usr/bin/env zsh
temporal-sql-tool --plugin $SQL_PLUGIN --ep $SQL_HOST -u $SQL_USER --pw $SQL_PASSWORD -p $DB_PORT --db $DBNAME setup-schema -v 0.0
temporal-sql-tool --plugin $SQL_PLUGIN --ep $SQL_HOST -u $SQL_USER --pw $SQL_PASSWORD -p $DB_PORT --db $DBNAME update-schema -d $PWD/../temporal/schema/postgresql/v96/temporal/versioned

temporal-sql-tool --plugin $VISIBILITY_SQL_PLUGIN --ep $VISIBILITY_SQL_HOST -u $VISIBILITY_SQL_USER --pw $VISIBILITY_SQL_PASSWORD -p $VISIBILITY_DB_PORT --db $VISIBILITY_DBNAME setup-schema -v 0.0
temporal-sql-tool --plugin $VISIBILITY_SQL_PLUGIN --ep $VISIBILITY_SQL_HOST -u $VISIBILITY_SQL_USER --pw $VISIBILITY_SQL_PASSWORD -p $VISIBILITY_DB_PORT --db $VISIBILITY_DBNAME update-schema -d $PWD/../temporal/schema/postgresql/v96/visibility/versioned