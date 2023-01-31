#!/usr/bin/env zsh
# assumes the directory structure:
# $parent
#   /distribution - this project
#   /temporal - upstream

SQL_PLUGIN=postgres
SQL_HOST=$(op item get nonprod-temporal-rds-us-west-2 --fields label=server)
SQL_USER=$(op item get nonprod-temporal-rds-us-west-2 --fields label=username)
SQL_PASSWORD=$(op item get nonprod-temporal-rds-us-west-2 --fields label=password)
SQL_PORT=$(op item get nonprod-temporal-rds-us-west-2 --fields label=port)
SQL_DATABASE=$(op item get nonprod-temporal-rds-us-west-2 --fields label=database)

temporal-sql-tool --plugin $SQL_PLUGIN --ep $SQL_HOST -u $SQL_USER --pw $SQL_PASSWORD -p $SQL_PORT --db $SQL_DATABASE setup-schema -v 0.0
temporal-sql-tool --plugin $SQL_PLUGIN --ep $SQL_HOST -u $SQL_USER --pw $SQL_PASSWORD -p $SQL_PORT --db $SQL_DATABASE update-schema -d $PWD/../temporal/schema/postgresql/v96/temporal/versioned

VISIBILITY_SQL_PLUGIN=postgres
VISIBILITY_SQL_HOST=$(op item get nonprod-temporal-visibility-rds-us-west-2 --fields label=server)
VISIBILITY_SQL_USER=$(op item get nonprod-temporal-visibility-rds-us-west-2 --fields label=username)
VISIBILITY_SQL_PASSWORD=$(op item get nonprod-temporal-visibility-rds-us-west-2 --fields label=password)
VISIBILITY_SQL_PORT=$(op item get nonprod-temporal-visibility-rds-us-west-2 --fields label=port)
VISIBILITY_SQL_DATABASE=$(op item get nonprod-temporal-visibility-rds-us-west-2 --fields label=database)

temporal-sql-tool --plugin $VISIBILITY_SQL_PLUGIN --ep $VISIBILITY_SQL_HOST -u $VISIBILITY_SQL_USER --pw $VISIBILITY_SQL_PASSWORD -p $VISIBILITY_SQL_PORT --db $VISIBILITY_SQL_DATABASE setup-schema -v 0.0
temporal-sql-tool --plugin $VISIBILITY_SQL_PLUGIN --ep $VISIBILITY_SQL_HOST -u $VISIBILITY_SQL_USER --pw $VISIBILITY_SQL_PASSWORD -p $VISIBILITY_SQL_PORT --db $VISIBILITY_SQL_DATABASE update-schema -d $PWD/../temporal/schema/postgresql/v96/visibility/versioned