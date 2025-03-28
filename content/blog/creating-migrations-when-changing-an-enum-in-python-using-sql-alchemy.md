---
title: Creating migrations when changing an enum in Python using SQLAlchemy
date: 2020-03-28 23:30:00
---
Recently at work I've spent some time working on a scaffold project that we'll be using for API projects that we'll be building.

For a number of reasons we've decided to use Python as our backend language and Flask as our API framework. One of the things I love most about Flask is that it's very unopinionated and let's you build what you want, pretty much how you want.

One of the features I've been working on is an audit log. For data integrity purposes we decided to use an enum field for the `event type` value in both the code itself and also in the database. Like many Flask applications, we're using SQLAlchemy as an ORM and Flask-Migrate to automatically create Alembic migrations. Using a SQLAlchemy model with the field type set to an enum equal to the enum used in the code, I had expected Flask-Migrate to automatically create a new migration any time we added values to the EventType enum class, however this is not the case.

After some searching around, I discovered that this is a known issue with Alembic and that migrations for enum changes have to be created manually. I've got a sample migration that I created manually below as well as some steps so you can see how to handle adding or removing a value from an enum in Flask. Note that this migration is specifically written to work with PostgreSQL as that is the database engine that we use.

1. Make changes to the Enum in the relevant model
2. Create an empty migration file ```flask db revision -m 'Add Logout_Success to AuditEvent'```
3. Populate the new migration with code to create the changes, note you will need to add values for the existing and new options ensuring to keep the revision and down_revision numbers that already exist in the new migration file

```pyhon
"""
Add Logout_Success to AuditEvent

Revision ID: 08720b8a9d11
Revises: 810eac468f83
Create Date: 2020-03-25 12:19:09.432635

"""
from alembic import op
import sqlalchemy as sa

# revision identifiers, used by Alembic.
revision = '08720b8a9d11'
down_revision = '810eac468f83'
branch_labels = None
depends_on = None

# Enum 'type' for PostgreSQL
enum_name = 'auditevent'
# Set temporary enum 'type' for PostgreSQL
tmp_enum_name = 'tmp_' + enum_name

# Options for Enum
old_options = ('LOGIN_SUCCESS', 'LOGIN_FAIL')
new_options = sorted(old_options + ('LOGOUT_SUCCESS',))

# Create enum fields
old_type = sa.Enum(*old_options, name=enum_name)
new_type = sa.Enum(*new_options, name=enum_name)

def upgrade():
    # Rename current enum type to tmp_
    op.execute('ALTER TYPE ' + enum_name + ' RENAME TO ' + tmp_enum_name)
    # Create new enum type in db
    new_type.create(op.get_bind())
    # Update column to use new enum type
    op.execute('ALTER TABLE audit ALTER COLUMN event_type TYPE ' + enum_name + ' USING event_type::text::' + enum_name)
    # Drop old enum type
    op.execute('DROP TYPE ' + tmp_enum_name)


def downgrade():
    # Instantiate db query
    audit = sa.sql.table('audit', sa.Column('event_type', new_type, nullable=False))
    # Convert LOGOUT_SUCCESS to LOGIN_SUCCESS (this is just a sample so may not make sense)
    op.execute(audit.update().where(audit.c.event_type == u'LOGOUT_SUCCESS').values(event_type='LOGIN_SUCCESS'))
    # Rename enum type to tmp_
    op.execute('ALTER TYPE ' + enum_name + ' RENAME TO ' + tmp_enum_name)
    # Create enum type using old values
    old_type.create(op.get_bind())
    # Set enum type as type for event_type column
    op.execute('ALTER TABLE audit ALTER COLUMN event_type TYPE ' + enum_name + ' USING event_type::text::' + enum_name)
    # Drop temp enum type
    op.execute('DROP TYPE ' + tmp_enum_name)

```
