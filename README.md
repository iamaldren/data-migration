# DATA MIGRATION

This is a sample program that will migrate records from Redis, which is in JSON format, to a typical table for Relational Databse.

## Work In Progress

Currently the code is just connecting to the database, and has a basic functions of Insert/Get statements.

## TODO

1. Get data from Redis (JSON format)
3. Call the functions that will execute statements in the database to migrate the data

## Additional Requirement

Since the relational database will implement an Auto-Increment (integer value) for the primary key, the Redis data must also adopt for this change.

Meaning the ID use in Redis will follow whatever is the new ID in database.
