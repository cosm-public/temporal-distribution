#!/usr/bin/env zsh
# assumes the directory structure:
# $parent
#   /distribution - this project
#   /temporal - upstream

temporal-sql-tool --plugin $SQL_PLUGIN --ep $SQL_HOST -u $SQL_USER --pw $SQL_PASSWORD -p $SQL_PORT --db $SQL_DATABASE setup-schema -v 0.0
temporal-sql-tool --plugin $SQL_PLUGIN --ep $SQL_HOST -u $SQL_USER --pw $SQL_PASSWORD -p $SQL_PORT --db $SQL_DATABASE update-schema -d $PWD/../temporal/schema/postgresql/v96/temporal/versioned

temporal-sql-tool --plugin $VISIBILITY_SQL_PLUGIN --ep $VISIBILITY_SQL_HOST -u $VISIBILITY_SQL_USER --pw $VISIBILITY_SQL_PASSWORD -p $VISIBILITY_SQL_PORT --db $VISIBILITY_SQL_DATABASE setup-schema -v 0.0
temporal-sql-tool --plugin $VISIBILITY_SQL_PLUGIN --ep $VISIBILITY_SQL_HOST -u $VISIBILITY_SQL_USER --pw $VISIBILITY_SQL_PASSWORD -p $VISIBILITY_SQL_PORT --db $VISIBILITY_SQL_DATABASE update-schema -d $PWD/../temporal/schema/postgresql/v96/visibility/versioned